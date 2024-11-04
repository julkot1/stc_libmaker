package main

import (
	"fmt"
	"os"
	"path/filepath"
	"stclibmake/config"
	"stclibmake/stc"
	"strconv"
)

func BuildLib(config *config.LibConfig, out string) {
	file, err := os.Create(filepath.Join(out, config.Head.Name+".c"))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writeHead(config.Head, config.Body.Method, file)
	writeBody(config.Body, file)
	createHeaderFile(config, out)

}

func createHeaderFile(cfg *config.LibConfig, out string) {
	file, err := os.Create(filepath.Join(out, cfg.Head.Name+".h"))

	if len(cfg.Head.Includes) != 0 {
		writeIncludes(cfg.Head.Includes, file)
	}

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	var methods []config.Method

	methods = append(methods, cfg.Head.Types.Method)
	methods = append(methods, cfg.Body.Method...)

	for _, method := range methods {
		name := method.Name
		if method.Stc {
			name = "stc_" + name
		}
		file.WriteString(method.Return + " " + name)
		writeArgs(method.Args, file)
		file.WriteString(";\n")
	}

}

func writeBody(body config.Body, file *os.File) {
	err := hasDuplicates(body.Method)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}

	for _, d := range body.Method {
		writeMethod(d, file)
	}

}

func writeMethod(method config.Method, file *os.File) {
	err := stc.IsValidCFunctionName(method.Name)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}
	rt := stc.ToCType(method.Return)

	err1 := stc.CheckReturnType(rt)
	if err1 != nil {
		fmt.Println("Error in method (return type) ", method.Name, ": ", err1)
		os.Exit(-1)
	}
	name := method.Name
	if method.Stc {
		name = "stc_" + name
	}
	file.WriteString("\n" + rt + " " + name)

	writeArgs(method.Args, file)

	file.WriteString("\n{\n")

	for _, line := range method.Code {
		file.WriteString("\t" + line + "\n")
	}

	file.WriteString("}\n")

}

func writeArgs(args []string, file *os.File) {
	file.WriteString("(")
	for i, arg := range args {
		argType := stc.ToCType(arg)
		_, err1 := stc.MatchTypeC(argType)
		if err1 != nil {
			fmt.Println("Error: ", err1)
			os.Exit(-1)
		}
		file.WriteString(argType + " arg" + strconv.Itoa(i))
		if i+1 != len(args) {
			file.WriteString(", ")
		}
	}
	file.WriteString(")")
}

func hasDuplicates(method []config.Method) error {
	seen := make(map[string]bool)

	for _, m := range method {
		str := getMethodSignature(m)
		if seen[str] == true {
			return fmt.Errorf("duplicate method: %s", m.Name)
		} else {
			seen[str] = true
		}
	}
	return nil
}

func getMethodSignature(m config.Method) string {
	str := m.Name + m.Return
	for _, arg := range m.Args {
		str = str + arg
	}
	return str
}

func writeHead(head config.Head, methods []config.Method, file *os.File) {

	file.WriteString("#include \"" + head.Name + ".h\"\n")
	if head.Types.Name != "" {
		writeTypes(head.Types, file)
		writeMatrix(head.Types, methods, file)
		writeMatchMethod(head.Types, methods, file)
	}

}

func writeMatchMethod(types config.HeadTypes, methods []config.Method, file *os.File) {
	if len(types.Args)*2 != len(types.Method.Args) {
		fmt.Println("Error: number of arguments does not match number of types in match method " + types.Name)
		os.Exit(-1)
	}
	matrixArgs := ""
	callArgs := ""

	for i, n := range types.Method.Args {
		if n == stc.C_TYPE_T.String() {
			matrixArgs = matrixArgs + "[arg" + strconv.Itoa(i) + "]"
		} else {
			callArgs = callArgs + "arg" + strconv.Itoa(i) + ","
		}
	}
	callArgs = callArgs[:len(callArgs)-1]

	types.Method.Code = append(types.Method.Code, types.TypeName+" f = "+types.Name+matrixArgs+";")
	types.Method.Code = append(types.Method.Code, "if(f != 0)")
	types.Method.Code = append(types.Method.Code, "\tf("+callArgs+")")

	writeMethod(types.Method, file)

}

func writeMatrix(types config.HeadTypes, methods []config.Method, file *os.File) {
	matrix := getMatrix(types.Match, methods, types)

	file.WriteString(types.TypeName + " " + types.Name + "[STC_TYPES_SIZE][STC_TYPES_SIZE] = {\n")
	for i := 0; i < len(matrix); i++ {
		file.WriteString("\t{")
		for j := 0; j < len(matrix[i]); j++ {
			file.WriteString(matrix[i][j])
			if j+1 != len(matrix) {
				file.WriteString(", ")
			}
		}
		if i+1 == len(matrix) {
			file.WriteString("}")
		} else {
			file.WriteString("},\n")
		}
	}
	file.WriteString("\n};\n")

}

func getMatrix(match []config.TypeMatch, methods []config.Method, types config.HeadTypes) [stc.SIZE_TYPE][stc.SIZE_TYPE]string {
	var matrix [stc.SIZE_TYPE][stc.SIZE_TYPE]string
	for i := 0; i < stc.SIZE_TYPE; i++ {
		for j := 0; j < stc.SIZE_TYPE; j++ {
			matrix[i][j] = "0"
		}
	}
	for _, t := range match {
		argA, err1 := stc.MatchTypeSTC(stc.ToSctType(t.ArgA))
		argB, err2 := stc.MatchTypeSTC(stc.ToSctType(t.ArgB))
		err3 := stc.ValidFunctionTypeMatrix(methods, t, types.Return, types.Args)
		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("Error:", err1, err2, err3)
			os.Exit(1)
		}
		method, err := stc.GetMethod(methods, t.Function)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		if method.Stc {
			matrix[argA][argB] = "stc_" + t.Function

		} else {
			matrix[argA][argB] = t.Function
		}
	}
	return matrix
}

func writeTypes(types config.HeadTypes, file *os.File) {

	args := getArgs(types.Args)
	file.WriteString("typedef " + types.Return + " (*" + types.TypeName + ")(" + args + ");\n")

}

func getArgs(args []string) string {
	str := ""
	for i, arg := range args {
		str = str + arg + " arg" + strconv.Itoa(i) + ", "
	}
	str = str[:len(str)-2]
	return str
}

func writeIncludes(includes []string, file *os.File) {
	for _, include := range includes {
		file.WriteString("#include \"" + include + "\"\n")
	}
}

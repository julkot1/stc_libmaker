package stc

import (
	"errors"
	"fmt"
	"stclibmake/config"
)

const (
	SIZE_TYPE  = 8
	SIZE_CTYPE = 3
)

type Type int64

const (
	INT_T Type = iota
	BOOL_T
	CHAR_T
	FLOAT_T
	STRING_T
	ARRAY_T
	STRUCT_T
	REF_T
)

type CType int64

const (
	C_I64_T CType = iota
	C_TYPE_T
	C_SIZE_T
	C_VOID_T
)

func (t Type) String() string {
	switch t {
	case INT_T:
		return "STC_I64_TYPE"
	case BOOL_T:
		return "STC_BOOL_TYPE"
	case CHAR_T:
		return "STC_CHAR_TYPE"
	case FLOAT_T:
		return "STC_FLOAT_TYPE"
	case STRING_T:
		return "STC_STRING_TYPE"
	case ARRAY_T:
		return "STC_ARRAY_TYPE"
	case STRUCT_T:
		return "STC_STRUCT_TYPE"
	case REF_T:
		return "STC_REF_TYPE"
	}
	return ""
}

func (t CType) String() string {
	switch t {
	case C_I64_T:
		return "STC_I64"
	case C_TYPE_T:
		return "STC_TYPE"
	case C_SIZE_T:
		return "STC_SIZE"
	}
	return "void"
}

func MatchTypeC(str string) (CType, error) {
	switch str {
	case C_I64_T.String():
		return C_I64_T, nil
	case C_TYPE_T.String():
		return C_TYPE_T, nil
	case C_SIZE_T.String():
		return C_SIZE_T, nil
	case "void":
		return C_VOID_T, nil
	}
	return -1, fmt.Errorf("invalid type: %s", str)
}

func MatchTypeSTC(str string) (Type, error) {
	switch str {
	case INT_T.String():
		return INT_T, nil
	case BOOL_T.String():
		return BOOL_T, nil
	case CHAR_T.String():
		return CHAR_T, nil
	case FLOAT_T.String():
		return FLOAT_T, nil
	case STRING_T.String():
		return STRING_T, nil
	case ARRAY_T.String():
		return ARRAY_T, nil
	case STRUCT_T.String():
		return STRUCT_T, nil
	case REF_T.String():
		return REF_T, nil
	default:
		return -1, errors.New("invalid type")
	}
}

func CheckReturnType(str string) error {
	if str == "void" {
		return nil
	}
	_, err := MatchTypeC(str)
	return err
}

func ToSctType(str string) string {
	return str + "_TYPE"
}

func ToCType(str string) string {
	if str == "void" {
		return "void"
	}
	return str

}

func FindMethod(methods []config.Method, name string) (config.Method, error) {
	for _, m := range methods {
		if m.Name == name {
			return m, nil
		}
	}
	return config.Method{}, errors.New("method " + name + " not found")
}

func ValidFunctionType(methods []config.Method, name string, returning string, args []string) error {
	method, err := FindMethod(methods, name)
	if err != nil {
		return err
	}
	if method.Return != returning {
		return fmt.Errorf("invalid returning type: %s", method.Return)
	}
	if len(args) != len(method.Args) {
		return fmt.Errorf("invalid number of arguments: %d", len(args))
	}
	for i, arg := range args {
		if arg != method.Args[i] {
			return fmt.Errorf("invalid argument type: %s", arg)
		}
	}
	return nil
}

func ValidFunctionTypeMatrix(methods []config.Method, match config.TypeMatch, returning string, args []string) error {
	return ValidFunctionType(methods, match.Function, returning, args)
}

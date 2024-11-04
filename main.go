package main

import (
	"fmt"
	"os"
	"path/filepath"
	"stclibmake/config"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <config>\n", os.Args[0])
		os.Exit(1)
	}
	path := os.Args[1]
	mainConf := config.LoadConfig(path)

	for _, lib := range mainConf.Libs {
		fmt.Println("Building Lib: " + lib)
		fullPath := filepath.Join(mainConf.RootInput, lib+".toml")
		c := config.LoadLibConfig(fullPath)
		BuildLib(c, mainConf.RootOutput)

	}
}

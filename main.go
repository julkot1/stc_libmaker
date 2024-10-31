package main

import (
	"stclibmake/config"
)

func main() {
	c := config.LoadLibConfig("arithmetic.toml")
	BuildLib(c)
}

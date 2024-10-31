package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

// LibConfig represents the full TOML structure
type LibConfig struct {
	Head Head `toml:"head"`
	Body Body `toml:"body"`
}

// Head represents the [head] section
type Head struct {
	Name     string    `toml:"name"`
	Includes []string  `toml:"includes"`
	Types    HeadTypes `toml:"types"`
}

// HeadTypes represents the [head.types] table within [head]
type HeadTypes struct {
	TypeName string      `toml:"type_name"`
	Name     string      `toml:"name"`
	Args     []string    `toml:"args"`
	Return   string      `toml:"return"`
	Method   Method      `toml:"method"`
	Match    []TypeMatch `toml:"match"`
}

// Method represents the [head.types.method] table within [head.types]
type Method struct {
	Name   string   `toml:"name"`
	Stc    bool     `toml:"stc"`
	Args   []string `toml:"args"`
	Return string   `toml:"return"`
	Code   []string `toml:"code"`
}

// TypeMatch represents each [[head.types.match]] entry
type TypeMatch struct {
	ArgA     string `toml:"argA"`
	ArgB     string `toml:"argB"`
	Function string `toml:"function"`
}

// Body represents the [body] section
type Body struct {
	Method []Method `toml:"method"`
}

func LoadLibConfig(path string) *LibConfig {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open TOML file: %v", err)
	}
	defer file.Close()

	config := &LibConfig{}
	if _, err := toml.NewDecoder(file).Decode(&config); err != nil {
		log.Fatalf("Failed to decode TOML: %v", err)
	}
	return config
}

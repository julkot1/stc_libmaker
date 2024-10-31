package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type LibsConfig struct {
	Root string   `toml:"root"`
	Libs []string `toml:"libs"`
}

func LoadConfig(path string) *LibsConfig {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open TOML file: %v", err)
	}
	defer file.Close()

	// Decode the TOML into the Config struct
	config := &LibsConfig{}
	if _, err := toml.NewDecoder(file).Decode(&config); err != nil {
		log.Fatalf("Failed to decode TOML file: %v", err)
	}
	return config

}

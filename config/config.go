package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

func LoadTOML(filename string, config any) error {
	_, err := toml.DecodeFile(filename, config)
	if err != nil {
		log.Printf("Error reading TOML file %s: %v", filename, err)
		return err
	}
	log.Printf("Successfully loaded TOML file: %s", filename)
	return nil
}

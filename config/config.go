package config

import (
	"log"
	"sync"
	"tasks/internal/models"

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

var (
	globalConfig *models.Config
	once         sync.Once
)

func LoadGlobalConfig(filename string) {
	once.Do(func() {
		var cfg models.Config
		if _, err := toml.DecodeFile(filename, &cfg); err != nil {
			log.Fatalf("Error loading config: %v", err)
		}
		globalConfig = &cfg
		log.Println("Global configuration loaded successfully")
	})
}

func GetConfig() *models.Config {
	if globalConfig == nil {
		log.Fatal("Configuration not loaded. Call LoadGlobalConfig first.")
	}
	return globalConfig
}

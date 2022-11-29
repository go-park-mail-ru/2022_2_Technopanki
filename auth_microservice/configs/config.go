package configs

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type SessionConfig struct {
	Domain                 string `yaml:"domain"`
	Port                   string `yaml:"port"`
	DefaultExpiringSession int    `yaml:"defaultExpiringSession"`
	GetConfirmInterval     int    `yaml:"getConfirmInterval"`
	ConfirmationTime       int    `yaml:"confirmationTime"`
	Redis                  RedisConfig
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

func InitConfig(config *SessionConfig) error {
	filename, fileErr := filepath.Abs("./auth_microservice/configs/config.yml")
	if fileErr != nil {
		return fileErr
	}

	yamlFile, yamlErr := os.ReadFile(filename)
	if yamlErr != nil {
		return yamlErr
	}

	marshalErr := yaml.Unmarshal(yamlFile, config)
	if marshalErr != nil {
		return marshalErr
	}

	return nil
}
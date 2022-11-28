package configs

import (
	"HeadHunter/configs"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type SessionConfig struct {
	DefaultExpiringSession int `yaml:"defaultExpiringSession"`
	ConfirmationTime       int `yaml:"confirmationTime"`
	Redis                  configs.RedisConfig
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

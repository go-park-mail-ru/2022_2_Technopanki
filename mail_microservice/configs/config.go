package configs

import (
	"HeadHunter/configs"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type MailConfig struct {
	Domain        string             `yaml:"domain"`
	Port          string             `yaml:"port"`
	Mail          configs.MailConfig `yaml:"mail"`
	SessionDomain string             `yaml:"sessionDomain"`
	SessionPort   string             `yaml:"sessionPort"`
}

func InitConfig(config *MailConfig) error {
	filename, fileErr := filepath.Abs("./mail_microservice/configs/config.yml")
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

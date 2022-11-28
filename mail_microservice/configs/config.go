package configs

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type Config struct {
	Domain     string     `yaml:"domain"`
	Port       string     `yaml:"port"`
	Mail       MailConfig `yaml:"mail"`
	AuthDomain string     `yaml:"authDomain"`
	AuthPort   string     `yaml:"authPort"`
}

type MailConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

func InitConfig(config *Config) error {
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

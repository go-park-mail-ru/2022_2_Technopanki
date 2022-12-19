package configs

import (
	"HeadHunter/configs"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type Config struct {
	Domain     string           `yaml:"domain"`
	Port       string           `yaml:"port"`
	Mail       MailConfig       `yaml:"mail"`
	AuthMsHost string           `yaml:"authMsHost"`
	AuthMsPort string           `yaml:"authMsPort"`
	MetricPath string           `yaml:"metricPath"`
	MetricPort string           `yaml:"metricPort"`
	DB         configs.DBConfig `yaml:"db"`
}

type MailConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

const dbPasswordName = "DB_PASSWORD"

func InitConfig(config *Config) error {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		return envErr
	}

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

	config.DB.Password = os.Getenv(dbPasswordName)
	return nil
}

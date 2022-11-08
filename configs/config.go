package configs

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type Config struct {
	Domain                 string           `yaml:"domain"`
	Port                   string           `yaml:"port"`
	DefaultExpiringSession int              `yaml:"defaultExpiringSession"`
	DB                     DBConfig         `yaml:"db"`
	Redis                  RedisConfig      `yaml:"redis"`
	Validation             ValidationConfig `yaml:"validation"`
	Cookie                 CookieConfig     `yaml:"cookie"`
	Crypt                  CryptConfig      `yaml:"crypt"`
	Image                  ImageConfig      `yaml:"image"`
	Security               SecurityConfig   `yaml:"security"`
}

type DBConfig struct {
	Username string `yaml:"username"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"sslmode"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type ValidationConfig struct {
	MinNameLength     int `yaml:"minNameLength"`
	MaxNameLength     int `yaml:"maxNameLength"`
	MinSurnameLength  int `yaml:"minSurnameLength"`
	MaxSurnameLength  int `yaml:"maxSurnameLength"`
	MinPasswordLength int `yaml:"minPasswordLength"`
	MaxPasswordLength int `yaml:"maxPasswordLength"`
	MinEmailLength    int `yaml:"minEmailLength"`
	MaxEmailLength    int `yaml:"maxEmailLength"`
}

type CookieConfig struct {
	HTTPOnly bool `yaml:"httpOnly"`
	Secure   bool `yaml:"secure"`
}

type CryptConfig struct {
	//COST The cost of the password encryption algorithm
	COST int `yaml:"cost"`
}

type ImageConfig struct {
	Path                   string `yaml:"path"`
	DefaultEmployerAvatar  string `yaml:"defaultEmployerAvatar"`
	DefaultApplicantAvatar string `yaml:"defaultApplicantAvatar"`
}

type SecurityConfig struct {
	Secret string `yaml:"csrfSecret"`
}

func InitConfig(config *Config) error {
	filename, fileErr := filepath.Abs("./configs/config.yml")
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

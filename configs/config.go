package configs

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type Config struct {
	Domain                 string           `yaml:"domain"`
	Port                   string           `yaml:"port"`
	AuthMsHost             string           `yaml:"authMsHost"`
	AuthMsPort             string           `yaml:"authMsPort"`
	MailMsHost             string           `yaml:"mailMsHost"`
	MailMsPort             string           `yaml:"mailMsPort"`
	MetricPath             string           `yaml:"metricPath"`
	CleaningPeriod         int64            `yaml:"cleaningPeriod"`
	DefaultExpiringSession int              `yaml:"defaultExpiringSession"`
	DB                     DBConfig         `yaml:"db"`
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

type ValidationConfig struct {
	MinNameLength              int `yaml:"minNameLength"`
	MaxNameLength              int `yaml:"maxNameLength"`
	MinSurnameLength           int `yaml:"minSurnameLength"`
	MaxSurnameLength           int `yaml:"maxSurnameLength"`
	MinPasswordLength          int `yaml:"minPasswordLength"`
	MaxPasswordLength          int `yaml:"maxPasswordLength"`
	MinEmailLength             int `yaml:"minEmailLength"`
	MaxEmailLength             int `yaml:"maxEmailLength"`
	MinResumeTitleLength       int `yaml:"minResumeTitleLength"`
	MaxResumeTitleLength       int `yaml:"maxResumeTitleLength"`
	MinResumeDescriptionLength int `yaml:"minResumeDescriptionLength"`
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
	Secret             string `yaml:"csrfSecret"`
	CsrfMode           bool   `yaml:"csrfMode"`
	ConfirmationTime   int    `yaml:"confirmationTime"`
	ConfirmAccountMode bool   `yaml:"confirmAccountMode"`
}

func InitConfig(config *Config) error {
	envErr := godotenv.Load("../.env")
	if envErr != nil {
		return envErr
	}

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

	config.DB.Password = os.Getenv("DB_PASSWORD")
	return nil
}

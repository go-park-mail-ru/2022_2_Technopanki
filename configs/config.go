package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	Domain                 string      `yaml:"domain"`
	Port                   string      `yaml:"port"`
	DefaultExpiringSession int64       `yaml:"defaultExpiringSession"`
	DB                     DBConfig    `yaml:"db"`
	Redis                  RedisConfig `yaml:"redis"`
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

func InitConfig(config *Config) error {
	//viper.AddConfigPath("configs")
	//viper.SetConfigName("config")
	//if err := viper.ReadInConfig(); err != nil {
	//	return err
	//}
	//config.Domain = viper.GetString("domain")
	//config.Port = viper.GetString("port")
	//config.DefaultExpiringSession = viper.GetInt64("defaultExpiringSession")
	//
	//config.DB.Host = viper.GetString("db.host")
	//config.DB.Port = viper.GetString("db.port")
	//config.DB.Username = viper.GetString("db.username")
	//config.DB.Password = viper.GetString("db.password")
	//config.DB.DBName = viper.GetString("db.dbname")
	//config.DB.SSLMode = viper.GetString("db.sslmode")
	//
	//config.Redis.Host = viper.GetString("redis.host")
	//config.Redis.Port = viper.GetString("redis.port")
	//config.Redis.Password = viper.GetString("redis.password")
	//config.Redis.DBName = viper.GetString("redis.dbname")

	filename, fileErr := filepath.Abs("./configs/config.yml")
	yamlFile, yamlErr := ioutil.ReadFile(filename)

	if fileErr != nil {
		return fileErr
	}
	if yamlErr != nil {
		return yamlErr
	}

	marshalErr := yaml.Unmarshal(yamlFile, config)
	if marshalErr != nil {
		return marshalErr
	}

	return nil
}

package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Domain string
	Port   string
	DB     DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func InitConfig(config *Config) error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	config.Domain = viper.GetString("domain")
	config.Port = viper.GetString("port")
	config.DB.Host = viper.GetString("db.host")
	config.DB.Port = viper.GetString("db.port")
	config.DB.Username = viper.GetString("db.username")
	config.DB.Password = viper.GetString("db.password")
	config.DB.DBName = viper.GetString("db.dbname")
	config.DB.SSLMode = viper.GetString("db.sslmode")
	return nil
}

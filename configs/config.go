package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Domain                 string
	Port                   string
	DefaultExpiringSession int64
	DB                     DBConfig
	Redis                  RedisConfig
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type RedisConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func InitConfig(config *Config) error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	config.Domain = viper.GetString("domain")
	config.Port = viper.GetString("port")
	config.DefaultExpiringSession = viper.GetInt64("defaultExpiringSession")

	config.DB.Host = viper.GetString("db.host")
	config.DB.Port = viper.GetString("db.port")
	config.DB.Username = viper.GetString("db.username")
	config.DB.Password = viper.GetString("db.password")
	config.DB.DBName = viper.GetString("db.dbname")
	config.DB.SSLMode = viper.GetString("db.sslmode")

	config.Redis.Host = viper.GetString("redis.host")
	config.Redis.Port = viper.GetString("redis.port")
	config.Redis.Password = viper.GetString("redis.password")
	config.Redis.DBName = viper.GetString("redis.dbname")
	return nil
}

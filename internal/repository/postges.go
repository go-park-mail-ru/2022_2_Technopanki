package repository

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {
	//dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	//	cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName)
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return &gorm.DB{}, errors.New("Cannot open" + err.Error())
	}

	pingErr := sqlDB.Ping()
	if pingErr != nil {
		return &gorm.DB{}, errors.New("Cannot ping" + pingErr.Error())

	}
	gormDB, gormErr := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if gormErr != nil {
		return &gorm.DB{}, gormErr
	}

	return gormDB, nil
}

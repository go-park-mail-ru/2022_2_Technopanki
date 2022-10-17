package repository

import (
	"HeadHunter/internal/entity"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func Connect(cfg DBConfig) error {
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//db, err := gorm.Open(postgres.Open("postgres://jobflowAdmin:12345@jfPostgres:5432/jobflowDB?sslmode=disable"), &gorm.Config{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&entity.User{}, &entity.Vacancy{}, &entity.Resume{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}

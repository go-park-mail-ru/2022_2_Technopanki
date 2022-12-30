package repository

import (
	"HeadHunter/internal/entity/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type DBConfig struct {
	Username           string `yaml:"username"`
	Host               string `yaml:"host"`
	Port               string `yaml:"port"`
	MaxIdleConnections int    `yaml:"maxConnections"`
	MaxOpenConnections int    `yaml:"maxOpenConnections"`
	DBName             string `yaml:"dbname"`
	Password           string `yaml:"password"`
	SSLMode            string `yaml:"sslmode"`
}

func DBConnect(cfg *DBConfig) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, sqlErr := db.DB()
	if sqlErr != nil {
		return nil, sqlErr
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnections)

	err = db.AutoMigrate(&models.UserAccount{}, &models.Resume{}, &models.EducationDetail{}, &models.ExperienceDetail{},
		&models.Vacancy{}, &models.Skill{}, &models.VacancyActivity{}, &models.BusinessType{}, &models.Notification{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

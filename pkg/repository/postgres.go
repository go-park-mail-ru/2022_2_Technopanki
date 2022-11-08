package repository

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnect(cfg configs.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.UserAccount{}, &models.Resume{}, &models.EducationDetail{}, &models.ExperienceDetail{},
		&models.Vacancy{}, &models.Skill{}, &models.VacancyActivity{}, &models.BusinessType{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

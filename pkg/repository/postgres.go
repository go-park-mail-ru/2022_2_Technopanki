package repository

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/Models"
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
	err = db.AutoMigrate(&Models.UserAccount{}, &Models.Resume{}, &Models.EducationDetail{}, &Models.ExperienceDetail{}, &Models.JobLocation{}, &Models.Vacancy{}, &Models.Skill{}, &Models.VacancyActivity{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

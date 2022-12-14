package cron

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"gorm.io/gorm"
	"time"
)

func ClearDBFromUnconfirmedUser(db *gorm.DB, cfg *configs.Config) {
	ticker := time.NewTicker(time.Duration(cfg.CleaningPeriod) * time.Hour)
	for range ticker.C {
		deleteUnconfirmedUsers(db)
	}
}

func deleteUnconfirmedUsers(db *gorm.DB) {
	user := &models.UserAccount{}
	db.Where("is_confirmed = false AND created_time < NOW() - '10 MINUTES'::INTERVAL").Delete(&user)
}

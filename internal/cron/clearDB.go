package cron

import (
	"HeadHunter/internal/entity/models"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func ClearDBFromUnconfirmedUser(db *gorm.DB, quit chan struct{}) {
	//ticker := time.NewTicker(2 * time.Hour)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-quit:
			return
		case <-ticker.C:
			fmt.Println("tick: ", time.Now().Second())
			deleteErr := deleteUnconfirmedUsers(db)
			if deleteErr != nil {
				fmt.Println(deleteErr)
			}
		default:
		}
	}
}

func deleteUnconfirmedUsers(db *gorm.DB) error {
	user := &models.UserAccount{}
	query := db.Where("is_confirmed = false AND created_time < NOW() - '10 MINUTES'::INTERVAL").Delete(&user)
	return query.Error
}

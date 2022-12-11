package impl

import (
	"HeadHunter/internal/entity/models"
	"fmt"
	"gorm.io/gorm"
)

type NotificationPostgres struct {
	db *gorm.DB
}

func NewNotificationPostgres(db *gorm.DB) *NotificationPostgres {
	return &NotificationPostgres{db: db}
}

func (np *NotificationPostgres) GetNotifications(id uint) ([]*models.Notification, error) {
	var result []*models.Notification
	query := np.db.Table("notifications").Where("user_to_id = ?", id).Scan(&result)
	return result, QueryValidation(query, "notification")

}

func (np *NotificationPostgres) CreateNotification(notification *models.Notification) error {
	creatingErr := np.db.Create(notification).Save(notification).Error

	if creatingErr != nil {
		return fmt.Errorf("cannot create notification: %w", creatingErr)
	}

	return nil
}

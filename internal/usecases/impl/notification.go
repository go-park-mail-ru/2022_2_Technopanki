package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/entity/utils"
	"HeadHunter/internal/repository"
	"HeadHunter/pkg/errorHandler"
	"errors"
)

type NotificationService struct {
	notificationRepo repository.NotificationRepository
	userRepo         repository.UserRepository
}

func NewNotificationService(notificationRepo repository.NotificationRepository, userRepo repository.UserRepository) *NotificationService {
	return &NotificationService{notificationRepo: notificationRepo, userRepo: userRepo}
}

func (ns *NotificationService) GetNotification(email string) ([]*models.Notification, error) {
	user, getErr := ns.userRepo.GetUserByEmail(email)
	if getErr != nil {
		return []*models.Notification{}, getErr
	}

	notifications, getErr := ns.notificationRepo.GetNotifications(user.ID)
	if errors.Is(getErr, errorHandler.ErrNotificationNotFound) {
		return []*models.Notification{}, nil
	}
	return notifications, getErr
}

func (ns *NotificationService) CreateNotification(notification *models.Notification) error {
	if !utils.HasStringArrayElement(notification.Type, models.AllowedNotificationTypes) {
		return errorHandler.ErrBadRequest
	}
	return ns.notificationRepo.CreateNotification(notification)
}

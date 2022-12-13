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

func (ns *NotificationService) GetNotificationsByEmail(email string) ([]*models.NotificationPreview, error) {
	user, getErr := ns.userRepo.GetUserByEmail(email)
	if getErr != nil {
		return []*models.NotificationPreview{}, getErr
	}
	var notifications []*models.NotificationPreview
	if user.UserType == "employer" {
		notifications, getErr = ns.notificationRepo.GetApplyNotificationsByUser(user.ID)
	} else {
		notifications, getErr = ns.notificationRepo.GetDownloadPDFNotificationsByUser(user.ID)
	}

	if errors.Is(getErr, errorHandler.ErrNotificationNotFound) {
		return []*models.NotificationPreview{}, nil
	}
	return notifications, getErr
}

func (ns *NotificationService) CreateNotification(notification *models.Notification) (*models.NotificationPreview, error) {
	if !utils.HasStringArrayElement(notification.Type, models.AllowedNotificationTypes) {
		return nil, errorHandler.ErrBadRequest
	}
	createErr := ns.notificationRepo.CreateNotification(notification)
	if createErr != nil {
		return nil, createErr
	}

	if notification.Type == "apply" {
		return ns.notificationRepo.GetNotificationPreviewApply(notification.ID)
	}
	return ns.notificationRepo.GetNotificationPreviewDownloadPDF(notification.ID)
}

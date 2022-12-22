package impl

import (
	"HeadHunter/internal/entity/models"
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
	if _, ok := models.AllowedNotificationTypes[notification.Type]; !ok {
		return nil, errorHandler.ErrBadRequest
	}

	if notification.UserFromID == notification.UserToID {
		return nil, nil
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

func (ns *NotificationService) ReadNotification(email string, id uint) error {
	notif, getNotifErr := ns.notificationRepo.GetNotification(id)

	if getNotifErr != nil {
		return getNotifErr
	}

	if notif.IsViewed {
		return errorHandler.ErrBadRequest
	}

	user, getUserErr := ns.userRepo.GetUserByEmail(email)

	if getUserErr != nil {
		return getUserErr
	}

	if user.ID != notif.UserToID {
		return errorHandler.ErrForbidden
	}

	return ns.notificationRepo.ReadNotification(id)
}

func (ns *NotificationService) ReadAllNotifications(email string) error {
	user, getUserErr := ns.userRepo.GetUserByEmail(email)

	if getUserErr != nil {
		return getUserErr
	}

	return ns.notificationRepo.ReadAllNotifications(user.ID)
}

func (ns *NotificationService) ClearNotifications(email string) error {
	user, getErr := ns.userRepo.GetUserByEmail(email)
	if getErr != nil {
		return getErr
	}

	return ns.notificationRepo.DeleteNotificationsFromUser(user.ID)
}

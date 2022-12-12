package impl

import (
	"HeadHunter/internal/network/handlers/utils"
	"HeadHunter/internal/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NotificationHandler struct {
	notificationUseCase usecases.Notification
}

func NewNotificationHandler(UseCase *usecases.UseCases) *NotificationHandler {
	return &NotificationHandler{notificationUseCase: UseCase.Notification}
}

func (nh *NotificationHandler) GetNotifications(c *gin.Context) {
	email, emailErr := utils.GetEmailFromContext(c)
	if emailErr != nil {
		_ = c.Error(emailErr)
		return
	}

	notifications, getErr := nh.notificationUseCase.GetNotificationsByEmail(email)
	if getErr != nil {
		_ = c.Error(getErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": notifications})
}

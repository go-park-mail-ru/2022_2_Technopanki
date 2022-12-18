package impl

import (
	"HeadHunter/internal/network/handlers/utils"
	"HeadHunter/internal/usecases"
	"HeadHunter/pkg/errorHandler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (nh *NotificationHandler) ReadNotification(c *gin.Context) {
	email, emailErr := utils.GetEmailFromContext(c)
	if emailErr != nil {
		_ = c.Error(emailErr)
		return
	}

	id, paramErr := strconv.Atoi(c.Param("id"))
	if paramErr != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	readErr := nh.notificationUseCase.ReadNotification(email, uint(id))
	if readErr != nil {
		_ = c.Error(readErr)
		return
	}

	c.Status(http.StatusOK)
}

func (nh *NotificationHandler) ReadAllNotifications(c *gin.Context) {
	email, emailErr := utils.GetEmailFromContext(c)
	if emailErr != nil {
		_ = c.Error(emailErr)
		return
	}

	readErr := nh.notificationUseCase.ReadAllNotifications(email)
	if readErr != nil {
		_ = c.Error(readErr)
		return
	}

	c.Status(http.StatusOK)
}

func (nh *NotificationHandler) ClearNotifications(c *gin.Context) {
	email, emailErr := utils.GetEmailFromContext(c)
	if emailErr != nil {
		_ = c.Error(emailErr)
		return
	}

	clearErr := nh.notificationUseCase.ClearNotifications(email)
	if clearErr != nil {
		_ = c.Error(clearErr)
		return
	}

	c.Status(http.StatusOK)
}

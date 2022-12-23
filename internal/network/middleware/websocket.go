package middleware

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/network/ws"
	"HeadHunter/pkg/errorHandler"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type WebSocketMiddleware struct {
	pool *ws.Pool
}

func NewWebsocketMiddleware(pool *ws.Pool) *WebSocketMiddleware {
	return &WebSocketMiddleware{pool: pool}
}

func (wsm *WebSocketMiddleware) Send(c *gin.Context) {
	if len(c.Errors) > 0 {
		return
	}

	notificationAny, okGet := c.Get("notification")
	notification, okCast := notificationAny.(*models.NotificationPreview)
	if !okGet || !okCast {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	if notification == nil {
		return
	}

	data, marshallErr := json.Marshal(notification)
	if marshallErr != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	getErr := wsm.pool.Send(notification.UserToID, data)
	if getErr != nil {
		_ = c.Error(getErr)
		return
	}
}

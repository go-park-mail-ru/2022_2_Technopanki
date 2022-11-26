package impl

import (
	"HeadHunter/internal/usecases"
	"HeadHunter/internal/usecases/mail"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MailHandler struct {
	mail mail.Mail
}

func NewMailHandler(cases *usecases.UseCases) *MailHandler {
	return &MailHandler{mail: cases.Mail}
}

func (mh *MailHandler) SendConfirmCode(c *gin.Context) {
	email := c.Param("email")
	confirmErr := mh.mail.SendConfirmCode(email)
	if confirmErr != nil {
		_ = c.Error(confirmErr)
		return
	}

	c.Status(http.StatusOK)
}

func (mh *MailHandler) UpdatePassword(c *gin.Context) {

}

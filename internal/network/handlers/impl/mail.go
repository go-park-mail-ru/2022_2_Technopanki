package impl

import (
	"HeadHunter/internal/usecases"
	"github.com/gin-gonic/gin"
)

type MailHandler struct {
	mail usecases.Mail
}

func NewMailHandler(cases *usecases.UseCases) *MailHandler {
	return &MailHandler{mail: cases.Mail}
}

func (mh *MailHandler) ConfirmationAccount(c *gin.Context) {

}
func (mh *MailHandler) UpdatePassword(c *gin.Context) {

}
func (mh *MailHandler) TwoFactorSignIn(c *gin.Context) {

}

package sender

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/mail_microservice/configs"
	"bytes"
	"errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/mail.v2"
	"html/template"
	"strings"
	"syscall"
)

type SenderService struct {
	dialer   *mail.Dialer
	username string
	cfg      *configs.Config
}

func NewSender(_cfg *configs.Config) (*SenderService, error) {

	dialer := mail.NewDialer(_cfg.Mail.Host, _cfg.Mail.Port, _cfg.Mail.Username, _cfg.Mail.Password)
	return &SenderService{dialer: dialer, username: _cfg.Mail.Username, cfg: _cfg}, nil
}

func (ss *SenderService) SendMail(to []string, subject, body string) error {
	msg := mail.NewMessage()
	msg.SetHeader("From", ss.username)
	msg.SetHeader("To", strings.Join(to, ", "))
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)
	sendErr := ss.dialer.DialAndSend(msg)
	if sendErr != nil {
		if errors.Is(sendErr, syscall.EPIPE) {
			logrus.Println("broken pipe")
			return ss.dialer.DialAndSend(msg)
		}
		return sendErr
	}
	return nil
}

func (ss *SenderService) SendConfirmCode(email, code string) error {
	form, parseErr := template.ParseFiles("./static/html/confirmLetter.html")
	if parseErr != nil {
		return parseErr
	}

	data := struct {
		Code string
	}{
		Code: code,
	}

	formBuf := bytes.NewBuffer([]byte(""))
	executeErr := form.Execute(formBuf, data)
	if executeErr != nil {
		return executeErr
	}

	return ss.SendMail([]string{email}, "Подтверждение аккаунта", formBuf.String())
}

func (ss *SenderService) SendApplicantMailing(email string, vacancies []*models.Vacancy) error {
	return nil
}

func (ss *SenderService) SendEmployerMailing(email string, applicants []*models.UserAccount) error {
	return nil
}

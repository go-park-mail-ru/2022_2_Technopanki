package sender

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"bytes"
	"errors"
	"fmt"
	"gopkg.in/mail.v2"
	"html/template"
	"strings"
	"syscall"
)

type SenderService struct {
	dial     mail.SendCloser
	username string
	cfg      *configs.Config
}

func NewSender(_cfg *configs.Config) (*SenderService, error) {

	dialer := mail.NewDialer(_cfg.Mail.Host, _cfg.Mail.Port, _cfg.Mail.Username, _cfg.Mail.Password)
	dial, dialErr := dialer.Dial()
	if dialErr != nil {
		return nil, dialErr
	}

	return &SenderService{dial: dial, username: _cfg.Mail.Username, cfg: _cfg}, nil
}

func (ss *SenderService) SendMail(to []string, subject, body string) error {
	msg := mail.NewMessage()
	msg.SetHeader("From", ss.username)
	msg.SetHeader("To", strings.Join(to, ", "))
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	sendErr := ss.dial.Send(ss.username, to, msg)
	if sendErr != nil {
		if errors.Is(sendErr, syscall.EPIPE) {
			fmt.Println("broken pipe")
			return nil
		}
		return sendErr
	}
	return nil
}

func (ss *SenderService) CloseSender() error {
	err := ss.dial.Close()
	if err != nil {
		return err
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

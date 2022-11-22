package sender

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"gopkg.in/gomail.v2"
	"strings"
)

type SenderService struct {
	dial     gomail.SendCloser
	username string
	cfg      *configs.Config
}

func NewSender(username, password string, _cfg *configs.Config) (*SenderService, error) {

	dialer := gomail.NewDialer("smtp.mail.ru", 587, username, password)
	dial, dialErr := dialer.Dial()
	if dialErr != nil {
		return nil, dialErr
	}

	return &SenderService{dial: dial, username: username, cfg: _cfg}, nil
}

func (ss *SenderService) SendMail(to []string, subject, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", ss.username)
	msg.SetHeader("To", strings.Join(to, ", "))
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	sendErr := ss.dial.Send(ss.username, to, msg)
	if sendErr != nil {
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

func (ss *SenderService) SendConfirmToken(email, token string) error {
	form := `<form action="http://` + ss.cfg.Domain + ss.cfg.Port + `/auth/confirm" method="post">
	 <div>
		<input type="hidden" name="token" value="` + token + `" />
	<button>Подтвердить аккаунт</button>
	 </div>
	</form>`
	return ss.SendMail([]string{email}, "Подтверждение аккаунта", form)
}

func (ss *SenderService) SendApplicantMailing(email string, vacancies []*models.Vacancy) error {
	return nil
}

func (ss *SenderService) SendEmployerMailing(email string, applicants []*models.UserAccount) error {
	return nil
}

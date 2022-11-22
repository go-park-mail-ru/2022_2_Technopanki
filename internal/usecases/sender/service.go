package sender

import (
	"HeadHunter/internal/entity/models"
	"gopkg.in/gomail.v2"
	"strings"
)

type SenderService struct {
	dial     gomail.SendCloser
	username string
}

func NewSender(username, password string) (*SenderService, error) {

	dialer := gomail.NewDialer("smtp.mail.ru", 587, username, password)
	dial, dialErr := dialer.Dial()
	if dialErr != nil {
		return nil, dialErr
	}

	return &SenderService{dial: dial, username: username}, nil
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

func (ss *SenderService) SendConfirmToken(token string) error {
	//	form := `<form action="http://localhost:8080" method="post">
	//  <div>
	//    <label for="say">What greeting do you want to say?</label>
	//    <input name="say" id="say" value="Hi">
	//  </div>
	//  <div>
	//	<label> your token is: ` + token + `</label>
	//    <button>Send my greetings</button>
	//  </div>
	//</form>`
	form := `<h1 style="color: blue"> your token is: ` + token + `</h1>`
	return ss.SendMail([]string{"zahvinar358@gmail.com"}, "Form", form)
}

func (ss *SenderService) SendApplicantMailing(email string, vacancies []*models.Vacancy) error {
	return nil
}

func (ss *SenderService) SendEmployerMailing(email string, applicants []*models.UserAccount) error {
	return nil
}

package cron

import (
	"HeadHunter/internal/repository"
	"HeadHunter/internal/usecases/mail"
	"github.com/sirupsen/logrus"
	"log"
)

func Mailing(repo repository.UserRepository, mail mail.Mail) func() {
	return func() {
		employerEmails, employerErr := repo.FindEmployersToMailing()
		if employerErr != nil {
			log.Println("errors with getting employers emails: ", employerErr)
		}

		resumes, resumeErr := repo.FindNewResumes()
		if resumeErr != nil {
			log.Println("errors with getting resumes: ", resumeErr)
		}

		err := mail.SendEmployerMailing(employerEmails, resumes)
		if err != nil {
			logrus.Println("send to applicant error: ", err)
		}

		applicantEmails, applicantErr := repo.FindApplicantsToMailing()
		if applicantErr != nil {
			log.Println("errors with getting applicants emails: ", applicantErr)
		}

		vacancies, vacancyErr := repo.FindNewVacancies()
		if vacancyErr != nil {
			log.Println("errors with getting vacancies: ", vacancyErr)
		}

		err = mail.SendApplicantMailing(applicantEmails, vacancies)
		if err != nil {
			logrus.Println("send to employer error: ", err)
		}
	}
}

package mail

type Mail interface {
	SendConfirmCode(email string) error
	SendApplicantMailing(email string) error
	SendEmployerMailing(email string) error
}

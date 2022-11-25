package mail

type Mail interface {
	SendConfirmCode(email string) error
	UpdatePassword()
}

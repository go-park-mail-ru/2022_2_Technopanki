package entity

type Applicant struct {
	ID       string `json:"-"`
	Name     string
	Surname  string
	Email    string
	Password string
	//Avatar
}

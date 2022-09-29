package entity

type Employer struct {
	ID       string `json:"-"`
	Name     string
	Email    string
	Password string
	//Logo
}

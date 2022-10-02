package entity

type User struct {
	ID       string `json:"-"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

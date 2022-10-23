package entity

type User struct {
	ID       string `json:"-" gorm:"primary_key"`
	Name     string `json:"name,omitempty"`
	Surname  string `json:"surname,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role,omitempty"`
}

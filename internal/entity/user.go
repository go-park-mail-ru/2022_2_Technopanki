package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `json:"id" gorm:"primary_key"`
	Name     string `json:"name,omitempty"`
	Surname  string `json:"surname,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role,omitempty"`
}

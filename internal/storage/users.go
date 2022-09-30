package storage

import (
	"HeadHunter/internal/entity"
)

type Users struct {
	// maps user email to entity.User
	Values map[string]entity.User
}

func (u *Users) IsUserInStorage(email string) bool {
	_, exists := u.Values[email]
	return exists
}

func (u *Users) AddUser(user entity.User) {
	u.Values[user.Email] = user
}

func NewUsersStorage() Users {
	return Users{
		Values: make(map[string]entity.User, 0),
	}
}

var UserStorage = NewUsersStorage()

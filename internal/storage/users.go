package storage

import (
	"HeadHunter/internal/entity"
	"HeadHunter/internal/errorHandler"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

//COST The cost of the password encryption algorithm
var COST = 10

type Users struct {
	// maps user email to entity.User
	Values map[string]entity.User
	mutex  sync.RWMutex
}

func (u *Users) GetUserByEmail(email string) (entity.User, error) {
	u.mutex.RLock()
	defer u.mutex.RUnlock()
	if val, exists := u.Values[email]; exists {
		return val, nil
	}

	return entity.User{}, errorHandler.ErrUserNotExists
}

func (u *Users) CreateUser(user entity.User) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), COST)
	if err != nil {
		return err
	}
	user.Password = string(encryptedPassword)

	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.Values[user.Email] = user

	return nil
}

var password, _ = bcrypt.GenerateFromPassword([]byte("123456!!a"), COST)

var UserStorage = Users{
	Values: map[string]entity.User{
		"example@mail.ru": {
			Name:     "Zakhar",
			Surname:  "Urvancev",
			Email:    "example@mail.ru",
			Password: string(password),
			Role:     "applicant",
		},
	},
}

package storage

import (
	"HeadHunter/internal/entity"
	"HeadHunter/internal/errorHandler"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

var COST = 10

type Users struct {
	// maps user email to entity.User
	Values map[string]entity.User
	mutex  sync.RWMutex
}

func (u *Users) FindByEmail(email string) (entity.User, error) {
	u.mutex.RLock()
	defer u.mutex.RUnlock()
	if val, exists := u.Values[email]; exists {
		return val, nil
	}

	return entity.User{}, errorHandler.ErrUserNotExists
}

func (u *Users) AddUser(user entity.User) error {
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

var UserStorage = Users{
	Values: make(map[string]entity.User, 0),
}

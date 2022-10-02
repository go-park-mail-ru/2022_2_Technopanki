package storage

import (
	"HeadHunter/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

type Users struct {
	// maps user email to entity.User
	Values map[string]entity.User
	mutex  sync.RWMutex
}

func (u *Users) FindByEmail(email string) entity.User {
	u.mutex.RLock()
	defer u.mutex.RUnlock()
	if val, exists := u.Values[email]; exists {
		return val
	}

	return entity.User{}
}

func (u *Users) AddUser(user entity.User) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(encryptedPassword)

	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.Values[user.Email] = user

	return nil
}

func NewUsersStorage() Users {
	return Users{
		Values: make(map[string]entity.User, 0),
	}
}

var UserStorage = NewUsersStorage()

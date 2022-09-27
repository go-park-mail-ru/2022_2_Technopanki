package storage

import (
	"errors"
	"main.go/internal/entities"
)

var Users = []entities.User{
	{
		ID:       1,
		Email:    "test@mail.ru",
		Name:     "Test",
		Password: "12345",
	},
}

func CreateUser(newUser entities.User) {
	newUser.ID = len(Users) + 1
	Users = append(Users, newUser)
}

func FindUserByEmail(email string) (entities.User, error) {
	for _, elem := range Users {
		if elem.Email == email {
			return elem, nil
		}
	}

	return entities.User{}, errors.New("user with this email doesnt exists")
}

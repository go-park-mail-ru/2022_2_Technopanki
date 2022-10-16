package usecases

import (
	"HeadHunter/internal/entity"
	"HeadHunter/internal/entity/validation"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/network/sessions"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/storage"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	SignUp(input entity.User) (string, error)
	SignIn(input *entity.User) (string, error)
	Logout(token string) error
	AuthCheck(email string) (entity.User, error)
}

type UserService struct {
	ur repository.UserRepository
}

func newUserService(userRepos repository.UserRepository) *UserService {
	return &UserService{ur: userRepos}
}

func (us *UserService) SignIn(input *entity.User) (string, error) {
	inputValidity := validation.IsValidateAuthData(*input)
	if inputValidity != nil {
		return "", inputValidity
	}
	user, err := storage.UserStorage.GetUserByEmail(input.Email)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", errorHandler.ErrUnauthorized
	}

	token := sessions.SessionsStore.NewSession(input.Email)
	input.Name = user.Name
	input.Surname = user.Surname
	return token, nil
}

func (us *UserService) SignUp(input entity.User) (string, error) {
	inputValidity := validation.IsValidate(input)
	if inputValidity != nil {
		return "", inputValidity
	}
	input.ID = uuid.NewString()
	_, err := storage.UserStorage.GetUserByEmail(input.Email)
	if err == nil {
		return "", errorHandler.ErrUserExists
	}
	err = storage.UserStorage.CreateUser(input)
	if err != nil {
		return "", errorHandler.ErrServiceUnavailable
	}
	token := sessions.SessionsStore.NewSession(input.Email)
	return token, us.ur.CreateUser(input)
}

func (us *UserService) Logout(token string) error {
	return sessions.SessionsStore.DeleteSession(sessions.Token(token))
}

func (us *UserService) AuthCheck(email string) (entity.User, error) {
	user, err := storage.UserStorage.GetUserByEmail(email)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

package usecases

import (
	"HeadHunter/internal/entity"
	"HeadHunter/internal/entity/validation"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/network/sessions"
	"HeadHunter/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	ur repository.UserRepository
}

func newUserService(userRepos repository.UserRepository) *UserService {
	return &UserService{ur: userRepos}
}

func (us *UserService) SignIn(input *entity.User) (string, error) {
	inputValidity := validation.IsAuthDataValid(*input)
	if inputValidity != nil {
		return "", inputValidity
	}
	user, err := us.ur.GetUserByEmail(input.Email)
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
	inputValidity := validation.IsUserValid(input)
	if inputValidity != nil {
		return "", inputValidity
	}
	_, err := us.ur.GetUserByEmail(input.Email)
	if err == nil {
		return "", errorHandler.ErrUserExists
	}
	err = us.ur.CreateUser(input)
	if err != nil {
		return "", errorHandler.ErrServiceUnavailable
	}
	input.ID = uuid.NewString()
	token := sessions.SessionsStore.NewSession(input.Email)
	return token, nil
}

func (us *UserService) Logout(token string) error {
	return sessions.SessionsStore.DeleteSession(sessions.Token(token))
}

func (us *UserService) AuthCheck(email string) (entity.User, error) {
	user, err := us.ur.GetUserByEmail(email)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

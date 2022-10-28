package usecases

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/entity/validation"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/repository/session"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRep repository.UserRepository
	sr      session.Repository
}

func newUserService(userRepos repository.UserRepository, sessionRepos session.Repository) *UserService {
	return &UserService{userRep: userRepos, sr: sessionRepos}
}

func (us *UserService) SignIn(input *models.UserAccount) (string, error) {
	inputValidity := validation.IsAuthDataValid(*input)
	if inputValidity != nil {
		return "", inputValidity
	}
	user, getErr := us.userRep.GetUserByEmail(input.Email)
	if getErr != nil {
		return "", getErr
	}
	if cryptErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); cryptErr != nil {
		return "", errorHandler.ErrUnauthorized
	}

	token, newSessionErr := us.sr.NewSession(input.Email)
	if newSessionErr != nil {
		return "", newSessionErr
	}
	if user.UserType == "applicant" {
		input.ApplicantName = user.ApplicantName
		input.ApplicantSurname = user.ApplicantSurname
	} else if user.UserType == "employer" {
		input.CompanyName = user.CompanyName
	} else {
		return "", errorHandler.InvalidUserType
	}
	return token, nil
}

func (us *UserService) SignUp(input *models.UserAccount) (string, error) {
	inputValidity := validation.IsUserValid(*input)
	if inputValidity != nil {
		return "", inputValidity
	}
	user, getErr := us.userRep.GetUserByEmail(input.Email)
	fmt.Println(user)
	if getErr == nil {
		return "", errorHandler.ErrUserExists
	}
	if getErr != nil && getErr != errorHandler.ErrUserNotExists {
		return "", getErr
	}

	createErr := us.userRep.CreateUser(*input)
	if createErr != nil {
		return "", errorHandler.ErrServiceUnavailable
	}
	input.UUID = uuid.NewString()

	token, newSessionErr := us.sr.NewSession(input.Email)
	if newSessionErr != nil {
		return "", newSessionErr
	}

	return token, nil
}

func (us *UserService) Logout(token string) error {
	return us.sr.DeleteSession(session.Token(token))
}

func (us *UserService) AuthCheck(email string) (models.UserAccount, error) {
	user, err := us.userRep.GetUserByEmail(email)
	if err != nil {
		return models.UserAccount{}, err
	}
	return *user, nil
}

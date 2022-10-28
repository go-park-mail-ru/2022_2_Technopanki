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
	user, err := us.userRep.GetUserByEmail(input.Email)
	if err != nil {
		return "", err
	}
	fmt.Println(user)
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
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
	user, err := us.userRep.GetUserByEmail(input.Email)
	fmt.Println(user)
	if err == nil {
		return "", errorHandler.ErrUserExists
	}
	if err != nil && err != errorHandler.ErrUserNotExists {
		return "", err
	}

	err = us.userRep.CreateUser(*input)
	if err != nil {
		return "", errorHandler.ErrServiceUnavailable
	}
	input.UUID = uuid.NewString()

	token, newSessionErr := us.sr.NewSession(input.Email)
	if newSessionErr != nil {
		return "", newSessionErr
	}

	//if input.UserType == "applicant" {
	//	input.ApplicantName = user.ApplicantName
	//	input.ApplicantSurname = user.ApplicantSurname
	//} else if input.UserType == "employer" {
	//	input.CompanyName = user.CompanyName
	//} else {
	//	return "", errorHandler.InvalidUserType
	//}

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

package usecases

import (
	"HeadHunter/internal/entity/Models"
	"HeadHunter/internal/entity/validation"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/repository/session"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	ur repository.UserRepository
	sr session.Repository
}

func newUserService(userRepos repository.UserRepository, sessionRepos session.Repository) *UserService {
	return &UserService{ur: userRepos, sr: sessionRepos}
}

func (us *UserService) SignIn(input *Models.UserAccount) (string, error) {
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

	token, newSessionErr := us.sr.NewSession(input.Email)
	if newSessionErr != nil {
		return "", newSessionErr
	}
	if input.UserType == "applicant" {
		input.ApplicantName = user.ApplicantName
		input.ApplicantSurname = user.ApplicantSurname
	} else if input.UserType == "employer" {
		input.CompanyName = user.CompanyName
	} else {
		return "", errorHandler.InvalidUserType
	}
	return token, nil
}

func (us *UserService) SignUp(input Models.UserAccount) (string, error) {
	inputValidity := validation.IsUserValid(input)
	if inputValidity != nil {
		return "", inputValidity
	}
	user, err := us.ur.GetUserByEmail(input.Email)
	if user != nil {
		return "", errorHandler.ErrUserExists
	}

	err = us.ur.CreateUser(input)
	if err != nil {
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

func (us *UserService) AuthCheck(email string) (Models.UserAccount, error) {
	user, err := us.ur.GetUserByEmail(email)
	if err != nil {
		return Models.UserAccount{}, err
	}
	return *user, nil
}

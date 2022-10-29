package usecases

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/entity/utils"
	"HeadHunter/internal/entity/validation"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/repository/session"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRep repository.UserRepository
	sr      session.Repository
	cfg     *configs.Config
}

func newUserService(userRepos repository.UserRepository, sessionRepos session.Repository, _cfg *configs.Config) *UserService {
	return &UserService{userRep: userRepos, sr: sessionRepos, cfg: _cfg}
}

func (us *UserService) SignIn(input *models.UserAccount) (string, error) {
	inputValidity := validation.IsAuthDataValid(input, us.cfg.Validation)
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

	if userCopyErr := utils.GetName(input, user); userCopyErr != nil {
		return "", userCopyErr
	}
	return token, nil
}

func (us *UserService) SignUp(input *models.UserAccount) (string, error) {
	inputValidity := validation.IsUserValid(input, us.cfg.Validation)
	if inputValidity != nil {
		return "", inputValidity
	}
	_, getErr := us.userRep.GetUserByEmail(input.Email)
	if getErr == nil {
		return "", errorHandler.ErrUserExists
	}
	if getErr != nil && getErr != errorHandler.ErrUserNotExists {
		return "", getErr
	}

	createErr := us.userRep.CreateUser(input)
	if createErr != nil {
		return "", errorHandler.ErrServiceUnavailable
	}

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

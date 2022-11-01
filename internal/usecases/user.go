package usecases

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/entity/utils"
	"HeadHunter/internal/entity/validation"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/repository/session"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	ur          repository.UserRepository
	sessionRepo session.Repository
	cfg         *configs.Config
}

func newUserService(userRepos repository.UserRepository, sessionRepos session.Repository, _cfg *configs.Config) *UserService {
	return &UserService{ur: userRepos, sessionRepo: sessionRepos, cfg: _cfg}
}

func (us *UserService) SignIn(input *models.UserAccount) (string, error) {
	inputValidity := validation.IsAuthDataValid(*input, us.cfg.Validation)
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

	token, newSessionErr := us.sessionRepo.NewSession(input.Email)
	if newSessionErr != nil {
		return "", newSessionErr
	}

	if userCopyErr := utils.FillUser(input, user); userCopyErr != nil {
		return "", userCopyErr
	}
	return token, nil
}

func (us *UserService) SignUp(input models.UserAccount) (string, error) {
	inputValidity := validation.IsUserValid(input, us.cfg.Validation)
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

	token, newSessionErr := us.sessionRepo.NewSession(input.Email)
	if newSessionErr != nil {
		return "", newSessionErr
	}
	return token, nil
}

func (us *UserService) Logout(token string) error {
	return us.sessionRepo.DeleteSession(token)
}

func (us *UserService) AuthCheck(email string) (models.UserAccount, error) {
	user, err := us.ur.GetUserByEmail(email)
	if err != nil {
		return models.UserAccount{}, err
	}
	return *user, nil
}

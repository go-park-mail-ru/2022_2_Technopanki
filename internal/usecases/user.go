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
  userRep     repository.UserRepository
  sessionRepo session.Repository
  cfg         *configs.Config
}

func newUserService(userRepos repository.UserRepository, sessionRepos session.Repository, _cfg *configs.Config) *UserService {
	return &UserService{userRep: userRepos, sessionRepo: sessionRepos, cfg: _cfg}
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
	if cryptErr := utils.ComparePassword(user, input); cryptErr != nil {
		return "", cryptErr
	}

	token, newSessionErr := us.sr.NewSession(input.Email)
	if newSessionErr != nil {
		return "", newSessionErr
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

func (us *UserService) SignUp(input *models.UserAccount) (string, error) {
	inputValidity := validation.IsUserValid(input, us.cfg.Validation)
	if inputValidity != nil {
		return "", inputValidity
	}
	isExist, getErr := us.userRep.IsUserExist(input.Email)
	if getErr != nil {
		return "", getErr
	}
	if isExist {
		return "", errorHandler.ErrUserExists
	}
	encryptedPassword, encryptErr := utils.GeneratePassword(input.Password, &us.cfg.Crypt)
	if encryptErr != nil {
		return "", encryptErr
	}
	input.Password = encryptedPassword
	createErr := us.userRep.CreateUser(input)
	if createErr != nil {
		return "", errorHandler.ErrCannotCreateUser
	}

	token, newSessionErr := us.sr.NewSession(input.Email)
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

func (us *UserService) UpgradeUser(input *models.UserAccount) error {
	inputValidity := validation.IsMainDataValid(input, us.cfg.Validation)
	if inputValidity != nil {
		return inputValidity
	}
	oldUser, getErr := us.userRep.GetUserByEmail(input.Email)
	if getErr != nil {
		return getErr
	} //TODO придумать более строгую валидацию
	dbError := us.userRep.UpgradeUser(oldUser, input)
	if dbError != nil {
		return dbError
	}
	return nil
}

func (us *UserService) GetUserId(email string) (uint, error) {
	user, getErr := us.userRep.GetUserByEmail(email)
	if getErr != nil {
		return 0, getErr
	}
	userId := user.ID
	return userId, nil
}

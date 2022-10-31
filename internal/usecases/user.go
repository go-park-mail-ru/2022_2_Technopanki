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
)

type UserService struct {
	userRep    repository.UserRepository
	sessionRep session.Repository
	cfg        *configs.Config
}

func newUserService(userRepos repository.UserRepository, sessionRepos session.Repository, _cfg *configs.Config) *UserService {
	return &UserService{userRep: userRepos, sessionRep: sessionRepos, cfg: _cfg}
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

	token, newSessionErr := us.sessionRep.NewSession(input.Email)
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
	input.Image = uuid.NewString()
	token, newSessionErr := us.sessionRep.NewSession(input.Email)
	if newSessionErr != nil {
		return "", newSessionErr
	}

	return token, nil
}

func (us *UserService) Logout(token string) error {
	return us.sessionRep.DeleteSession(session.Token(token))
}

func (us *UserService) AuthCheck(email string) (*models.UserAccount, error) {
	user, err := us.userRep.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) UpdateUser(input *models.UserAccount) error {
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

func (us *UserService) GetUser(id uint) (*models.UserAccount, error) {
	return us.userRep.GetUser(id)
}

func (us *UserService) GetUserSafety(id uint) (*models.UserAccount, error) {
	safeFields := []string{"email", "user_type", "contact_number", "description", "date_of_birth",
		"applicant_name", "applicant_surname", "company_name", "applicant_current_salary",
		"business_type", "company_website_url", "resumes", "vacancies"}
	return us.userRep.GetUserSafety(id, safeFields)
}

func (us *UserService) GetUserImage(id uint) {

}

func (us *UserService) UpdateUserImage() {

}

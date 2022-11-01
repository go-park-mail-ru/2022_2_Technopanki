package usecases

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/entity/utils"
	"HeadHunter/internal/entity/validation"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/repository/images"
	"HeadHunter/internal/repository/session"
	"fmt"
	"github.com/google/uuid"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"mime/multipart"
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
	input.Image = fmt.Sprintf("basic_%s_avatar.webp", input.UserType)
	createErr := us.userRep.CreateUser(input)
	if createErr != nil {
		return "", fmt.Errorf("creating session user: %w", createErr)
	}
	token, newSessionErr := us.sessionRepo.NewSession(input.Email)
	if newSessionErr != nil {
		return "", newSessionErr
	}
	return token, nil
}

func (us *UserService) Logout(token string) error {
	return us.sessionRepo.DeleteSession(token)
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
	}

	var encryptErr error = nil
	var encryptedPassword string
	if input.Password != "" {
		encryptedPassword, encryptErr = utils.GeneratePassword(input.Password, &us.cfg.Crypt)
	}
	if encryptErr != nil {
		return encryptErr
	}
	input.Password = encryptedPassword

	dbError := us.userRep.UpdateUser(oldUser, input)
	if dbError != nil {
		return dbError
	}
	return nil
}

func (us *UserService) GetUser(id uint) (*models.UserAccount, error) {
	return us.userRep.GetUser(id)
}

func (us *UserService) GetUserSafety(id uint) (*models.UserAccount, error) {
	return us.userRep.GetUserSafety(id, []string{}) //TODO добавить поле в бд
}

func (us *UserService) GetUserByEmail(email string) (*models.UserAccount, error) {
	return us.userRep.GetUserByEmail(email)
}

func (us *UserService) UploadUserImage(user *models.UserAccount, file *multipart.File, imageExt string) error {
	var img image.Image
	var decodeErr error
	switch imageExt {
	case "jpeg":
		img, decodeErr = jpeg.Decode(*file)
	case "jpg":
		img, decodeErr = jpeg.Decode(*file)
	case "png":
		img, decodeErr = png.Decode(*file)
	case "gif":
		img, decodeErr = gif.Decode(*file)
	default:
		decodeErr = errorHandler.ErrInvalidFileFormat
	}
	if decodeErr != nil {
		return decodeErr
	}

	if user.Image == fmt.Sprintf("basic_%s_avatar.webp", user.UserType) || user.Image == "" {
		user.Image = fmt.Sprintf("%s.webp", uuid.NewString())
	}

	return images.UploadUserAvatar(user.Image, &img, &us.cfg.Image)

}

func (us *UserService) DeleteUserImage(user *models.UserAccount) error {
	if user.Image == fmt.Sprintf("basic_%s_avatar.webp", user.UserType) || user.Image == "" {
		return errorHandler.ErrBadRequest
	}
	deleteErr := images.DeleteUserAvatar(user.Image, &us.cfg.Image)
	if deleteErr != nil {
		return deleteErr
	}
	user.Image = fmt.Sprintf("basic_%s_avatar.webp", user.UserType)
	return nil
}

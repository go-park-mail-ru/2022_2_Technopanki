package impl

import (
	"HeadHunter/common/session"
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/entity/utils"
	"HeadHunter/internal/entity/validation"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/repository/images"
	"HeadHunter/internal/usecases/escaping"
	"HeadHunter/internal/usecases/mail"
	"HeadHunter/pkg/errorHandler"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"reflect"
	"strings"
	"time"
)

type UserService struct {
	userRep     repository.UserRepository
	sessionRepo session.Repository
	cfg         *configs.Config
	mail        mail.Mail
}

func NewUserService(userRepos repository.UserRepository, sessionRepos session.Repository,
	_mail mail.Mail, _cfg *configs.Config) *UserService {
	return &UserService{userRep: userRepos, sessionRepo: sessionRepos, mail: _mail, cfg: _cfg}
}

func (us *UserService) GetUserId(email string) (uint, error) {
	user, getErr := us.GetUserByEmail(email)
	if getErr != nil {
		return 0, getErr
	}
	return user.ID, nil
}

func (us *UserService) SignIn(input *models.UserAccount) (string, error) {
	inputValidity := validation.IsAuthDataValid(input, us.cfg.Validation)
	if inputValidity != nil {
		return "", inputValidity
	}

	input = escaping.EscapingObject[*models.UserAccount](input)

	user, getErr := us.userRep.GetUserByEmail(input.Email)
	if getErr != nil {
		return "", getErr
	}

	if !user.IsConfirmed && us.cfg.Security.ConfirmAccountMode {
		return "", errorHandler.ErrIsNotConfirmed
	}

	if cryptErr := utils.ComparePassword(user, input); cryptErr != nil {
		return "", cryptErr
	}

	if user.TwoFactorSignIn {
		sendErr := us.mail.SendConfirmCode(user.Email)
		if sendErr != nil {
			return "", sendErr
		}
		return "", errorHandler.ErrForbidden
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

	input = escaping.EscapingObject[*models.UserAccount](input)

	user, getErr := us.userRep.GetUserByEmail(input.Email)

	if getErr == nil {
		if us.cfg.Security.ConfirmAccountMode {
			if user.IsConfirmed {
				return "", errorHandler.ErrUserExists
			}
			_, getCodeErr := us.sessionRepo.GetSession(input.Email)
			if getCodeErr == nil {
				return "", errorHandler.ErrIsNotConfirmed
			}

			sendCodeErr := us.mail.SendConfirmCode(input.Email)
			if sendCodeErr != nil {
				return "", sendCodeErr
			}
			return "", errorHandler.ErrIsNotConfirmed
		}
		return "", errorHandler.ErrUserExists
	}

	if getErr != errorHandler.ErrUserNotExists {
		return "", getErr
	}

	if us.cfg.Security.ConfirmAccountMode {
		sendCodeErr := us.mail.SendConfirmCode(input.Email)
		if sendCodeErr != nil {
			return "", sendCodeErr
		}
	}

	encryptedPassword, encryptErr := utils.GeneratePassword(input.Password, &us.cfg.Crypt)
	if encryptErr != nil {
		return "", encryptErr
	}
	input.Password = encryptedPassword

	input.Image = fmt.Sprintf("basic_%s_avatar.webp", input.UserType)
	input.PublicFields = "email contact_number applicant_current_salary" //TODO после РК3 убрать для добавления фичи с доступом
	input.IsConfirmed = !us.cfg.Security.ConfirmAccountMode

	input.Age = uint(Age(input.DateOfBirth))

	createErr := us.userRep.CreateUser(input)

	if createErr != nil {
		return "", fmt.Errorf("creating session user: %w", createErr)
	}

	var token string
	var newSessionErr error
	if !us.cfg.Security.ConfirmAccountMode {
		token, newSessionErr = us.sessionRepo.NewSession(input.Email)
		if newSessionErr != nil {
			return "", newSessionErr
		}
	}
	return token, nil
}

func (us *UserService) Logout(token string) error {
	return us.sessionRepo.Delete(token)
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

	input = escaping.EscapingObject[*models.UserAccount](input)

	oldUser, getErr := us.userRep.GetUserByEmail(input.Email)
	if getErr != nil {
		return getErr
	}

	if oldUser.UserType != input.UserType {
		return errorHandler.ErrBadRequest
	}

	if input.Password != "" {
		encryptedPassword, encryptErr := utils.GeneratePassword(input.Password, &us.cfg.Crypt)
		if encryptErr != nil {
			return encryptErr
		}

		input.Password = encryptedPassword
	}
	input.ID = oldUser.ID
	if input.Image == "" {
		input.Image = oldUser.Image
	}
	input.IsConfirmed = oldUser.IsConfirmed

	input.Age = uint(Age(input.DateOfBirth))

	dbError := us.userRep.UpdateUser(input)
	if dbError != nil {
		return dbError
	}

	return nil
}

func (us *UserService) GetUser(id uint) (*models.UserAccount, error) {
	return us.userRep.GetUser(id)
}

func (us *UserService) GetAllEmployers(filters models.UserFilter) ([]*models.UserAccount, error) {
	var conditions []string
	var filterValues []interface{}
	values := reflect.ValueOf(filters)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).Interface().(string) != "" {
			query := EmployerFilterQueries(types.Field(i).Name)
			if query != "" {
				conditions = append(conditions, query)
			}
			filterValues = append(filterValues, values.Field(i).Interface().(string))
		}
	}
	return us.userRep.GetAllUsers(conditions, filterValues, "employer")
}

func (us *UserService) GetAllApplicants(filters models.UserFilter) ([]*models.UserAccount, error) {
	var conditions []string
	var filterValues []interface{}
	values := reflect.ValueOf(filters)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).Interface().(string) != "" {
			query := ApplicantFilterQueries(types.Field(i).Name)
			if query != "" {
				conditions = append(conditions, query)
			}
			filterValues = append(filterValues, values.Field(i).Interface().(string))
		}
	}
	return us.userRep.GetAllUsers(conditions, filterValues, "applicant")
}

func (us *UserService) GetUserSafety(id uint) (*models.UserAccount, error) {
	user, getErr := us.userRep.GetUser(id)
	if getErr != nil {
		return nil, getErr
	}

	if validErr := validation.AllowedFieldsValidation(user); validErr != nil {
		return nil, validErr
	}
	fields := strings.Split(user.PublicFields, " ")

	if len(fields) == 1 && (fields[0] == "" || fields[0] == models.NoPublicFields) {
		fields = []string{}
	}
	return us.userRep.GetUserSafety(id, fields)
}

func (us *UserService) GetUserByEmail(email string) (*models.UserAccount, error) {
	return us.userRep.GetUserByEmail(email)
}

func (us *UserService) UploadUserImage(user *models.UserAccount, fileHeader *multipart.FileHeader) (string, error) {
	file, fileErr := fileHeader.Open()
	if fileErr != nil {
		return "", fileErr
	}

	user = escaping.EscapingObject[*models.UserAccount](user)
	imageName := fmt.Sprintf("%d.webp", user.ID)

	user.Image = fmt.Sprintf("%d.webp?%s", user.ID, time.Now().String())

	updateErr := us.userRep.UpdateUser(&models.UserAccount{ID: user.ID, Image: user.Image})
	if updateErr != nil {
		return "", updateErr
	}

	img, _, decodeErr := image.Decode(file)
	if decodeErr != nil {
		fmt.Println("Error in decoding (UploadUserImage)")
		return "", errorHandler.ErrBadRequest
	}

	return user.Image, images.UploadUserAvatar(imageName, &img, &us.cfg.Image)
}

func (us *UserService) DeleteUserImage(user *models.UserAccount) error {
	user = escaping.EscapingObject[*models.UserAccount](user)

	if user.Image == fmt.Sprintf("basic_%s_avatar.webp", user.UserType) || user.Image == "" {
		return errorHandler.ErrBadRequest
	}
	deleteErr := images.DeleteUserAvatar(user.Image, &us.cfg.Image)
	if deleteErr != nil {
		return errorHandler.ErrCannotDeleteAvatar
	}
	user.Image = fmt.Sprintf("basic_%s_avatar.webp", user.UserType)
	return us.UpdateUser(user)
}

func (us *UserService) ConfirmUser(code, email string) (*models.UserAccount, string, error) {
	if code == "" {
		return nil, "", errorHandler.ErrBadRequest
	}

	codeFromEmail, getCodeErr := us.sessionRepo.GetCodeFromEmail(email)
	if getCodeErr != nil {
		return nil, "", getCodeErr
	}

	if code != codeFromEmail {
		return nil, "", errorHandler.ErrForbidden
	}

	user, getErr := us.userRep.GetUserByEmail(email)
	if getErr != nil {
		return nil, "", getErr
	}

	confirmedUser := user
	confirmedUser.IsConfirmed = true

	token, newSessionErr := us.sessionRepo.NewSession(email)
	if newSessionErr != nil {
		return nil, "", newSessionErr
	}

	return confirmedUser, token, us.userRep.UpdateUser(confirmedUser)
}

func (us *UserService) UpdatePassword(code, email, password string) error {
	codeFromCode, getCodeErr := us.sessionRepo.GetCodeFromEmail(email)
	if getCodeErr != nil {
		return getCodeErr
	}

	if code != codeFromCode {
		return errorHandler.ErrForbidden
	}

	user, getErr := us.userRep.GetUserByEmail(email)
	if getErr != nil {
		return getErr
	}

	encryptedPassword, encryptErr := utils.GeneratePassword(password, &us.cfg.Crypt)
	if encryptErr != nil {
		return encryptErr
	}
	user.Password = encryptedPassword
	return us.userRep.UpdateUser(user)
}

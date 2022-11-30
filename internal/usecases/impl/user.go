package impl

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/entity/utils"
	"HeadHunter/internal/entity/validation"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/repository/images"
	"HeadHunter/internal/repository/session"
	"HeadHunter/internal/usecases/escaping"
	"HeadHunter/pkg/errorHandler"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"reflect"
	"strings"
)

type UserService struct {
	userRep     repository.UserRepository
	sessionRepo session.Repository
	cfg         *configs.Config
}

func NewUserService(userRepos repository.UserRepository, sessionRepos session.Repository, _cfg *configs.Config) *UserService {
	return &UserService{userRep: userRepos, sessionRepo: sessionRepos, cfg: _cfg}
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

	input = escaping.EscapingObject[*models.UserAccount](input)

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
	input.PublicFields = "email contact_number applicant_current_salary" //TODO после РК3 убрать для добавления фичи с доступом
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
	dbError := us.userRep.UpdateUser(oldUser, input)
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

	if user.Image == fmt.Sprintf("basic_%s_avatar.webp", user.UserType) || user.Image == "" {
		user.Image = fmt.Sprintf("%d.webp", user.ID)

		updateErr := us.UpdateUser(user)
		if updateErr != nil {
			return "", updateErr
		}
	}

	img, _, decodeErr := image.Decode(file)
	if decodeErr != nil {
		fmt.Println("Error in decoding (UploadUserImage)")
		return "", errorHandler.ErrBadRequest
	}

	return user.Image, images.UploadUserAvatar(user.Image, &img, &us.cfg.Image)
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

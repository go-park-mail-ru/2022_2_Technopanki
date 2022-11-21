package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/pkg/errorHandler"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (up *UserPostgres) CreateUser(user *models.UserAccount) error {
	return up.db.Create(user).Error
}

func (up *UserPostgres) UpdateUser(oldUser, newUser *models.UserAccount) error {
	return up.db.Model(oldUser).Updates(newUser).Error
}
func (up *UserPostgres) UpdateUserField(oldUser, newUser *models.UserAccount, field ...string) error {
	return up.db.Model(oldUser).Select(field).Updates(newUser).Error
}
func (up *UserPostgres) GetUserByEmail(email string) (*models.UserAccount, error) {
	var result models.UserAccount
	query := up.db.Where("email = ?", email).Find(&result)
	return &result, QueryValidation(query, "user")
}

func (up *UserPostgres) IsUserExist(email string) (bool, error) {
	_, getErr := up.GetUserByEmail(email)
	if getErr == nil {
		return true, nil
	}
	if getErr == errorHandler.ErrUserNotExists {
		return false, nil
	}
	return false, getErr
}

func (up *UserPostgres) GetUser(id uint) (*models.UserAccount, error) {
	var result models.UserAccount
	query := up.db.Select(append(models.PrivateUserFields, models.SafeUserFields...)).Find(&result, id)
	return &result, QueryValidation(query, "user")
}

func (up *UserPostgres) GetUserSafety(id uint, allowedFields []string) (*models.UserAccount, error) {
	var result models.UserAccount

	query := up.db.Select(append(models.SafeUserFields, allowedFields...)).Find(&result, id)
	return &result, QueryValidation(query, "user")
}

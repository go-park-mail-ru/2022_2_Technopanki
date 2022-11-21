package impl

import (
	"HeadHunter/internal/entity/models"
	"gorm.io/gorm"
)

type AuthorPostgres struct {
	db *gorm.DB
}

func NewAuthorPostgres(db *gorm.DB) *AuthorPostgres {
	return &AuthorPostgres{db: db}
}

func (ap *AuthorPostgres) GetAuthor(email string) (*models.UserAccount, error) {
	var result models.UserAccount
	query := ap.db.Where("email = ?", email).Find(&result)
	return &result, QueryValidation(query, "user")
}

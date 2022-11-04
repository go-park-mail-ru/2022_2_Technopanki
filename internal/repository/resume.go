package repository

import (
	"HeadHunter/internal/entity/models"
	"fmt"
	"gorm.io/gorm"
)

type ResumePostgres struct {
	db *gorm.DB
}

func newResumePostgres(db *gorm.DB) *ResumePostgres {
	return &ResumePostgres{db: db}
}

func (rp *ResumePostgres) Get(id uint) (*models.Resume, error) {
	var result models.Resume
	query := rp.db.First(&result, id)
	return &result, queryValidation(query, "resume")
}

func (rp *ResumePostgres) GetByApplicant(userId uint) ([]*models.Resume, error) {
	var result []*models.Resume
	query := rp.db.Where("user_account_id = ?", userId).Find(&result)
	return result, queryValidation(query, "resume")
}

func (rp *ResumePostgres) Create(resume *models.Resume, userId uint) (uint, error) {
	resume.UserAccountId = userId
	creatingErr := rp.db.Create(resume).Error
	if creatingErr != nil {
		return 0, fmt.Errorf("cannot create resume: %w", creatingErr)
	}
	return resume.ID, nil
}

func (rp *ResumePostgres) Update(id uint, resume *models.Resume) error {
	return rp.db.Model(&models.Resume{ID: id}).Updates(resume).Error
}

func (rp *ResumePostgres) Delete(id uint) error {
	return rp.db.Delete(&models.Resume{ID: id}).Error
}

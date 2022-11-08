package repository

import (
	"HeadHunter/internal/entity/models"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type ResumePostgres struct {
	db *gorm.DB
}

func newResumePostgres(db *gorm.DB) *ResumePostgres {
	return &ResumePostgres{db: db}
}

func (rp *ResumePostgres) GetResume(id uint) (*models.Resume, error) {
	var result models.Resume
	query := rp.db.First(&result, id)

	rp.db.Where("resume_id = ?", result.ID).First(&result.ExperienceDetail)

	rp.db.Where("resume_id = ?", result.ID).First(&result.EducationDetail)

	return &result, queryValidation(query, "resume")
}

func (rp *ResumePostgres) GetResumeByApplicant(userId uint) ([]*models.Resume, error) {
	var result []*models.Resume
	query := rp.db.Where("user_account_id = ?", userId).Find(&result)

	for _, elem := range result {
		rp.db.Where("resume_id = ?", elem.ID).Find(&elem.ExperienceDetail)

		rp.db.Where("resume_id = ?", elem.ID).Find(&elem.EducationDetail)
	}

	return result, queryValidation(query, "resume")
}

func (rp *ResumePostgres) CreateResume(resume *models.Resume, userId uint) error {
	resume.UserAccountId = userId
	creatingErr := rp.db.Create(resume).Save(resume).Error
	if creatingErr != nil {
		return fmt.Errorf("cannot create resume: %w", creatingErr)
	}
	resume.CreatedTime = time.Now()
	return nil
}

func (rp *ResumePostgres) UpdateResume(id uint, resume *models.Resume) error {
	old, getErr := rp.GetResume(id)
	fmt.Println(old)
	if getErr != nil {
		return getErr
	}
	resume.UserAccountId = old.UserAccountId
	resume.ID = id
	return rp.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(resume).Error
}

func (rp *ResumePostgres) DeleteResume(id uint) error {
	return rp.db.Delete(&models.Resume{ID: id}).Error
}

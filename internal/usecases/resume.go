package usecases

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/usecases/sanitize"
)

type ResumeService struct {
	resumeRep repository.ResumeRepository
	cfg       *configs.Config
}

func newResumeService(_resumeRep repository.ResumeRepository, _cfg *configs.Config) *ResumeService {
	return &ResumeService{resumeRep: _resumeRep, cfg: _cfg}
}

func (rs *ResumeService) GetResume(id uint) (*models.Resume, error) {
	return rs.resumeRep.GetResume(id)
}

func (rs *ResumeService) GetResumeByApplicant(userId uint) ([]*models.Resume, error) {
	return rs.resumeRep.GetResumeByApplicant(userId)
}

func (rs *ResumeService) GetPreviewResumeByApplicant(userId uint) ([]*models.Resume, error) {
	return rs.resumeRep.GetPreviewResumeByApplicant(userId)
}

func (rs *ResumeService) CreateResume(resume *models.Resume, userId uint) error {
	var sanitizeErr error
	resume, sanitizeErr = sanitize.SanitizeObject[*models.Resume](resume)
	if sanitizeErr != nil {
		return sanitizeErr
	}

	return rs.resumeRep.CreateResume(resume, userId)
}

func (rs *ResumeService) UpdateResume(id uint, resume *models.Resume) error {
	var sanitizeErr error
	resume, sanitizeErr = sanitize.SanitizeObject[*models.Resume](resume)
	if sanitizeErr != nil {
		return sanitizeErr
	}

	return rs.resumeRep.UpdateResume(id, resume)
}

func (rs *ResumeService) DeleteResume(id uint) error {
	return rs.resumeRep.DeleteResume(id)
}

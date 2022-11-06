package usecases

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/repository"
)

type ResumeService struct {
	resumeRep repository.ResumeRepository
	cfg       *configs.Config
}

func newResumeService(_resumeRep repository.ResumeRepository, _cfg *configs.Config) *ResumeService {
	return &ResumeService{resumeRep: _resumeRep, cfg: _cfg}
}

func (rs *ResumeService) Get(id uint) (*models.Resume, error) {
	return rs.resumeRep.Get(id)
}

func (rs *ResumeService) GetByApplicant(userId uint) ([]*models.Resume, error) {
	return rs.resumeRep.GetByApplicant(userId)
}

func (rs *ResumeService) Create(resume *models.Resume, userId uint) (uint, error) {
	return rs.resumeRep.Create(resume, userId)
}

func (rs *ResumeService) Update(id uint, resume *models.Resume) error {
	return rs.resumeRep.Update(id, resume)
}

func (rs *ResumeService) Delete(id uint) error {
	return rs.resumeRep.Delete(id)
}

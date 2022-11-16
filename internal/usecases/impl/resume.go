package impl

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/entity/validation"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/usecases"
	"HeadHunter/internal/usecases/sanitize"
)

type ResumeService struct {
	resumeRep repository.ResumeRepository
	cfg       *configs.Config
}

func NewResumeService(_resumeRep repository.ResumeRepository, _cfg *configs.Config) *ResumeService {
	return &ResumeService{resumeRep: _resumeRep, cfg: _cfg}
}

func (rs *ResumeService) GetResume(id uint) (*models.Resume, error) {
	return rs.resumeRep.GetResume(id)
}

func (rs *ResumeService) GetResumeByApplicant(userId uint, email string) ([]*models.Resume, error) {
	userFromContext, contextErr := usecases.GetUser(email, rs.resumeRep.GetDB())
	if contextErr != nil {
		return []*models.Resume{}, contextErr
	}

	if userFromContext.ID != userId {
		return []*models.Resume{}, errorHandler.ErrUnauthorized
	}
	return rs.resumeRep.GetResumeByApplicant(userId)
}

func (rs *ResumeService) GetPreviewResumeByApplicant(userId uint, email string) ([]*models.ResumePreview, error) {
	userFromContext, contextErr := usecases.GetUser(email, rs.resumeRep.GetDB())
	if contextErr != nil {
		return []*models.ResumePreview{}, contextErr
	}

	if userFromContext.ID != userId {
		return []*models.ResumePreview{}, errorHandler.ErrUnauthorized
	}
	return rs.resumeRep.GetPreviewResumeByApplicant(userId)
}

func (rs *ResumeService) CreateResume(resume *models.Resume, email string) error {

	isResumeValid := validation.ResumeValidaion(resume, rs.cfg.Validation)
	if isResumeValid != nil {
		return isResumeValid
	}

	user, getErr := usecases.GetUser(email, rs.resumeRep.GetDB())
	if getErr != nil {
		return getErr
	}

	if user.UserType != "applicant" {
		return errorHandler.InvalidUserType
	}

	var sanitizeErr error
	resume, sanitizeErr = sanitize.SanitizeObject[*models.Resume](resume)
	if sanitizeErr != nil {
		return sanitizeErr
	}

	return rs.resumeRep.CreateResume(resume, user.ID)
}

func (rs *ResumeService) UpdateResume(id uint, resume *models.Resume, email string) error {

	isResumeValid := validation.ResumeValidaion(resume, rs.cfg.Validation)
	if isResumeValid != nil {
		return isResumeValid
	}

	userFromContext, contextErr := usecases.GetUser(email, rs.resumeRep.GetDB())
	if contextErr != nil {
		return contextErr
	}

	old, getErr := rs.resumeRep.GetResume(id)
	if getErr != nil {
		return getErr
	}
	resume.UserAccountId = old.UserAccountId
	resume.ID = id

	if userFromContext.ID != old.UserAccountId {
		return errorHandler.ErrUnauthorized
	}

	var sanitizeErr error
	resume, sanitizeErr = sanitize.SanitizeObject[*models.Resume](resume)
	if sanitizeErr != nil {
		return sanitizeErr
	}

	return rs.resumeRep.UpdateResume(id, resume)
}

func (rs *ResumeService) DeleteResume(id uint, email string) error {
	userFromContext, contextErr := usecases.GetUser(email, rs.resumeRep.GetDB())
	if contextErr != nil {
		return contextErr
	}

	old, getErr := rs.resumeRep.GetResume(id)
	if getErr != nil {
		return getErr
	}

	if userFromContext.ID != old.UserAccountId {
		return errorHandler.ErrUnauthorized
	}
	return rs.resumeRep.DeleteResume(id)
}

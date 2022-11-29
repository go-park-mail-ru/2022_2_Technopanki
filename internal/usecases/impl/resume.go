package impl

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/entity/validation"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/usecases/escaping"
	"HeadHunter/pkg/errorHandler"
	"errors"
)

type ResumeService struct {
	resumeRep repository.ResumeRepository
	cfg       *configs.Config
	userRep   repository.UserRepository
}

func NewResumeService(_resumeRep repository.ResumeRepository, _cfg *configs.Config, _userRep repository.UserRepository) *ResumeService {
	return &ResumeService{resumeRep: _resumeRep, cfg: _cfg, userRep: _userRep}
}

func (rs *ResumeService) GetResume(id uint, email string) (*models.Resume, error) {
	userFromContext, contextErr := rs.userRep.GetUserByEmail(email)
	if contextErr != nil {
		return nil, contextErr
	}

	resume, getErr := rs.resumeRep.GetResume(id)
	if getErr != nil {
		return nil, getErr
	}

	if userFromContext.ID != resume.UserAccountId {
		employerId, err := rs.resumeRep.GetEmployerIdByVacancyActivity(id)
		if err != nil || employerId != userFromContext.ID {
			return nil, errorHandler.ErrUnauthorized
		}
	}

	return resume, nil
}

func (rs *ResumeService) GetResumeByApplicant(userId uint, email string) ([]*models.Resume, error) {
	//userFromContext, contextErr := rs.userRep.GetUserByEmail(email)
	//if contextErr != nil {
	//	return []*models.Resume{}, contextErr
	//}
	//
	//if userFromContext.ID != userId {
	//	return []*models.Resume{}, errorHandler.ErrUnauthorized
	//}
	resumes, getErr := rs.resumeRep.GetResumeByApplicant(userId)
	if errors.Is(getErr, errorHandler.ErrResumeNotFound) {
		return []*models.Resume{}, nil
	}
	return resumes, getErr
}

func (rs *ResumeService) GetPreviewResumeByApplicant(userId uint, email string) ([]*models.ResumePreview, error) {
	//userFromContext, contextErr := rs.userRep.GetUserByEmail(email)
	//if contextErr != nil {
	//	return []*models.ResumePreview{}, contextErr
	//}
	//
	//if userFromContext.ID != userId {
	//	return []*models.ResumePreview{}, errorHandler.ErrUnauthorized
	//}
	resumesPreview, getErr := rs.resumeRep.GetPreviewResumeByApplicant(userId)
	if errors.Is(getErr, errorHandler.ErrResumeNotFound) {
		return []*models.ResumePreview{}, nil
	}
	return resumesPreview, getErr
}

func (rs *ResumeService) CreateResume(resume *models.Resume, email string) error {

	isResumeValid := validation.ResumeValidaion(resume, rs.cfg.Validation)
	if isResumeValid != nil {
		return isResumeValid
	}

	user, getErr := rs.userRep.GetUserByEmail(email)
	if getErr != nil {
		return getErr
	}

	if user.UserType != "applicant" {
		return errorHandler.InvalidUserType
	}

	resume = escaping.EscapingObject[*models.Resume](resume)

	return rs.resumeRep.CreateResume(resume, user.ID)
}

func (rs *ResumeService) UpdateResume(id uint, resume *models.Resume, email string) error {

	isResumeValid := validation.ResumeValidaion(resume, rs.cfg.Validation)
	if isResumeValid != nil {
		return isResumeValid
	}

	userFromContext, contextErr := rs.userRep.GetUserByEmail(email)
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

	resume = escaping.EscapingObject[*models.Resume](resume)

	return rs.resumeRep.UpdateResume(id, resume)
}

func (rs *ResumeService) DeleteResume(id uint, email string) error {
	userFromContext, contextErr := rs.userRep.GetUserByEmail(email)
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

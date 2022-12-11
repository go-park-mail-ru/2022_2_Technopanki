package impl

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/entity/validation"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/usecases/escaping"
	"HeadHunter/internal/usecases/utils"
	"HeadHunter/pkg/errorHandler"
	"errors"
	"reflect"
)

type ResumeService struct {
	resumeRep repository.ResumeRepository
	cfg       *configs.Config
	userRep   repository.UserRepository
}

func NewResumeService(_resumeRep repository.ResumeRepository, _cfg *configs.Config, _userRep repository.UserRepository) *ResumeService {
	return &ResumeService{resumeRep: _resumeRep, cfg: _cfg, userRep: _userRep}
}

func (rs *ResumeService) GetResume(id uint) (*models.Resume, error) {

	resume, getErr := rs.resumeRep.GetResume(id)
	if getErr != nil {
		return nil, getErr
	}

	return resume, nil
}

func (rs *ResumeService) GetAllResumes(filters models.ResumeFilter) ([]*models.Resume, error) {
	var conditions []string
	var filterValues []interface{}
	values := reflect.ValueOf(filters)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).Interface().(string) != "" {
			query := ResumeFilterQueries(types.Field(i).Name)
			if query != "" {
				conditions = append(conditions, query)
			}
			filterValues = append(filterValues, values.Field(i).Interface().(string))
		}
	}
	return rs.resumeRep.GetAllResumes(conditions, filterValues)
}

func (rs *ResumeService) GetResumeByApplicant(userId uint) ([]*models.Resume, error) {
	resumes, getErr := rs.resumeRep.GetResumeByApplicant(userId)
	if errors.Is(getErr, errorHandler.ErrResumeNotFound) {
		return []*models.Resume{}, nil
	}
	return resumes, getErr
}

func (rs *ResumeService) GetPreviewResumeByApplicant(userId uint) ([]*models.ResumePreview, error) {
	resumesPreview, getErr := rs.resumeRep.GetPreviewResumeByApplicant(userId)
	if errors.Is(getErr, errorHandler.ErrResumeNotFound) {
		return []*models.ResumePreview{}, nil
	}
	return resumesPreview, getErr
}

func (rs *ResumeService) GetResumeInPDF(resumeId uint) ([]byte, error) {
	resumeInPDFModel, getErr := rs.resumeRep.GetResumeInPDF(resumeId)
	if getErr != nil {
		return nil, getErr
	}

	resumeInPDF, generateErr := utils.GenerateResumeInPDF(resumeInPDFModel)
	if generateErr != nil {
		return nil, generateErr
	}

	return resumeInPDF, nil
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

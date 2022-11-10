package validation

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
)

func ResumeValidaion(resume *models.Resume, cfg configs.ValidationConfig) error {
	if len(resume.Title) > cfg.MaxResumeTitleLength || len(resume.Title) < cfg.MinResumeTitleLength {
		return errorHandler.InvalidResumeTitleLength
	}
	if len(resume.Description) < cfg.MinResumeDescriptionLength {
		return errorHandler.InvalidResumeDescriptionLength
	}
	if resume.UserSurname != "" || resume.UserName != "" || resume.ImgSrc != "" {
		return errorHandler.ErrBadRequest
	}
	return nil
}

package validation

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/pkg/errorHandler"
)

func ResumeValidaion(resume *models.Resume, cfg configs.ValidationConfig) error {
	if len([]rune(resume.Title)) > cfg.MaxResumeTitleLength || len([]rune(resume.Title)) < cfg.MinResumeTitleLength {
		return errorHandler.InvalidResumeTitleLength
	}
	if len([]rune(resume.Description)) < cfg.MinResumeDescriptionLength {
		return errorHandler.InvalidResumeDescriptionLength
	}
	return nil
}

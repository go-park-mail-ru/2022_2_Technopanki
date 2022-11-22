package utils

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/pkg/errorHandler"
)

func FillUser(user, reference *models.UserAccount) error {
	user.ID = reference.ID
	user.UserType = reference.UserType
	user.Image = reference.Image
	if reference.UserType == "applicant" {
		user.ApplicantName = reference.ApplicantName
		user.ApplicantSurname = reference.ApplicantSurname
	} else if reference.UserType == "employer" {
		user.CompanyName = reference.CompanyName
	} else {
		return errorHandler.InvalidUserType
	}
	return nil
}

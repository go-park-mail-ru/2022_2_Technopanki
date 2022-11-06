package utils

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
)

func FillUser(user, reference *models.UserAccount) error {
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

//func GetUserId(c *gin.Context) (uint, error) {
//	userEmail, ok := c.Get("userEmail")
//	if !ok {
//		paramErr := c.Error(errorHandler.ErrInvalidParam)
//		return 0, paramErr
//	}
//	emailString := userEmail.(string)
//	us := usecases.UserService{}
//	userId, userIdErr := us.GetUserId(emailString)
//	if userIdErr != nil {
//		err := c.Error(userIdErr)
//		return 0, err
//	}
//	return userId, nil
//}

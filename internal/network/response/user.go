package response

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func generateUserResponse(user *models.UserAccount, fields []string) (models.UserAccount, error) {
	result := models.UserAccount{}

	for _, field := range fields {
		switch field {
		case "id":
			result.ID = user.ID
		case "user_type":
			if user.UserType != "applicant" && user.UserType != "employer" {
				return models.UserAccount{}, errorHandler.InvalidUserType
			}
			result.UserType = user.UserType
		case "image":
			result.Image = user.Image
		case "email":
			result.Email = user.Email
		case "status":
			result.Status = user.Status
		case "name_data":
			if user.UserType == "applicant" {
				result.ApplicantName = user.ApplicantName
				result.ApplicantSurname = user.ApplicantSurname
			} else if user.UserType == "employer" {
				result.CompanyName = user.CompanyName
			} else {
				return models.UserAccount{}, errorHandler.InvalidUserType
			}
		}
	}
	return result, nil
}

func sendUserResponse(c *gin.Context, user *models.UserAccount, fields []string) {
	result, generateErr := generateUserResponse(user, fields)
	if generateErr != nil {
		_ = c.Error(generateErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func SendSuccessData(c *gin.Context, user *models.UserAccount) {
	sendUserResponse(c, user, []string{"id", "user_type", "name_data", "image"})
}

func SendPreviewData(c *gin.Context, user *models.UserAccount) {
	sendUserResponse(c, user, []string{"id", "user_type", "name_data", "image", "status"})
}

func SendUploadImageData(c *gin.Context, user *models.UserAccount) {
	sendUserResponse(c, user, []string{"image"})
}

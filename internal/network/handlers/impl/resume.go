package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/network/handlers/utils"
	"HeadHunter/internal/usecases"
	"HeadHunter/pkg/errorHandler"
	"HeadHunter/pkg/themes"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type ResumeHandler struct {
	resumeUseCase usecases.Resume
}

func NewResumeHandler(useCases *usecases.UseCases) *ResumeHandler {
	return &ResumeHandler{resumeUseCase: useCases.Resume}
}

func (rh *ResumeHandler) GetResume(c *gin.Context) {
	id, paramErr := strconv.Atoi(c.Param("id"))
	if paramErr != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	resume, getResumeErr := rh.resumeUseCase.GetResume(uint(id))
	if getResumeErr != nil {
		_ = c.Error(getResumeErr)
		return
	}

	c.JSON(http.StatusOK, resume)
}

func (rh *ResumeHandler) GetAllResumes(c *gin.Context) {
	var resumes []*models.Resume
	var getAllErr error
	var filters models.ResumeFilter

	title := c.Query("search")
	if title != "" {
		filters.Title = title
	}

	experience := c.Query("experience")
	if experience != "" {
		filters.ExperienceInYears = experience
	}

	city := c.Query("city")
	if city != "" {
		filters.Location = city
	}

	salary := c.Query("salary")
	if salary != "" {
		if strings.Contains(salary, ":") {
			split := strings.Split(salary, ":")
			filters.FirstSalaryValue = split[0]
			filters.SecondSalaryValue = split[1]
		} else {
			_ = c.Error(errorHandler.ErrBadRequest)
			return
		}
	}

	resumes, getAllErr = rh.resumeUseCase.GetAllResumes(filters)
	if getAllErr != nil {
		_ = c.Error(getAllErr)
		return
	}
	c.JSON(http.StatusOK, models.GetAllResumesResponcePointer{
		Data: resumes,
	})
}

func (rh *ResumeHandler) GetResumeByApplicant(c *gin.Context) {
	userId, paramErr := strconv.Atoi(c.Param("user_id"))
	if paramErr != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	resumes, getResumeErr := rh.resumeUseCase.GetResumeByApplicant(uint(userId))
	if getResumeErr != nil {
		_ = c.Error(getResumeErr)
		return
	}

	c.JSON(http.StatusOK, resumes)
}

func (rh *ResumeHandler) GetPreviewResumeByApplicant(c *gin.Context) {
	userId, paramErr := strconv.Atoi(c.Param("user_id"))
	if paramErr != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	resumes, getResumeErr := rh.resumeUseCase.GetPreviewResumeByApplicant(uint(userId))
	if getResumeErr != nil {
		_ = c.Error(getResumeErr)
		return
	}

	c.JSON(http.StatusOK, resumes)
}

func (rh *ResumeHandler) GetResumeInPDF(c *gin.Context) {
	id, idErr := strconv.Atoi(c.Param("id"))
	if idErr != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	style := c.Query("style")
	zoomSize, err := strconv.ParseFloat(c.Query("zoom_size"), 64)
	if err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	resumeStyle, exists := themes.ThemesMap[style]
	if !exists {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	resumeInPDF, generateErr := rh.resumeUseCase.GetResumeInPDF(uint(id), resumeStyle, zoomSize)
	if generateErr != nil {
		_ = c.Error(generateErr)
		return
	}

	c.Data(http.StatusOK, "application/pdf", resumeInPDF)
}

func (rh *ResumeHandler) CreateResume(c *gin.Context) {

	email, contextErr := utils.GetEmailFromContext(c)
	if contextErr != nil {
		_ = c.Error(contextErr)
		return
	}

	var input models.Resume
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	creatingErr := rh.resumeUseCase.CreateResume(&input, email)
	if creatingErr != nil {
		_ = c.Error(creatingErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": input.ID})
}

func (rh *ResumeHandler) UpdateResume(c *gin.Context) {
	id, paramErr := strconv.Atoi(c.Param("id"))
	if paramErr != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	email, contextErr := utils.GetEmailFromContext(c)
	if contextErr != nil {
		_ = c.Error(contextErr)
		return
	}

	var input models.Resume
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	updateErr := rh.resumeUseCase.UpdateResume(uint(id), &input, email)
	if updateErr != nil {
		_ = c.Error(updateErr)
		return
	}

	c.Status(http.StatusOK)
}

func (rh *ResumeHandler) DeleteResume(c *gin.Context) {
	id, paramErr := strconv.Atoi(c.Param("id"))
	if paramErr != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	email, contextErr := utils.GetEmailFromContext(c)
	if contextErr != nil {
		_ = c.Error(contextErr)
		return
	}

	deleteErr := rh.resumeUseCase.DeleteResume(uint(id), email)
	if deleteErr != nil {
		_ = c.Error(deleteErr)
		return
	}

	c.Status(http.StatusOK)
}

package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/usecases"
	"HeadHunter/pkg/errorHandler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	email, contextErr := getEmailFromContext(c)
	if contextErr != nil {
		_ = c.Error(contextErr)
		return
	}

	resume, getResumeErr := rh.resumeUseCase.GetResume(uint(id), email)
	if getResumeErr != nil {
		_ = c.Error(getResumeErr)
		return
	}

	c.JSON(http.StatusOK, resume)
}

func (rh *ResumeHandler) GetResumeByApplicant(c *gin.Context) {
	userId, paramErr := strconv.Atoi(c.Param("user_id"))
	if paramErr != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	email, contextErr := getEmailFromContext(c)
	if contextErr != nil {
		_ = c.Error(contextErr)
		return
	}

	resumes, getResumeErr := rh.resumeUseCase.GetResumeByApplicant(uint(userId), email)
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

	email, contextErr := getEmailFromContext(c)
	if contextErr != nil {
		_ = c.Error(contextErr)
		return
	}

	resumes, getResumeErr := rh.resumeUseCase.GetPreviewResumeByApplicant(uint(userId), email)
	if getResumeErr != nil {
		_ = c.Error(getResumeErr)
		return
	}

	c.JSON(http.StatusOK, resumes)
}

func (rh *ResumeHandler) CreateResume(c *gin.Context) {

	email, contextErr := getEmailFromContext(c)
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

	email, contextErr := getEmailFromContext(c)
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

	email, contextErr := getEmailFromContext(c)
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

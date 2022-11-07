package handlers

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ResumeHandler struct {
	cfg           *configs.Config
	resumeUseCase usecases.Resume
	userHandler   UserH
}

func newResumeHandler(useCases *usecases.UseCases, _cfg *configs.Config, _userHandler UserH) *ResumeHandler {
	return &ResumeHandler{cfg: _cfg, resumeUseCase: useCases.Resume, userHandler: _userHandler}
}

func (rh *ResumeHandler) GetResume(c *gin.Context) {
	id, paramErr := strconv.Atoi(c.Param("id"))
	if paramErr != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	isAccessAllowed := rh.isResumeAvailable(c, uint(id))
	if isAccessAllowed != nil {
		_ = c.Error(isAccessAllowed)
		return
	}

	resume, getResumeErr := rh.resumeUseCase.GetResume(uint(id))
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

	userIdFromContext, getUserErr := rh.userHandler.GetUserId(c)
	if getUserErr != nil {
		_ = c.Error(getUserErr)
		return
	}

	if userIdFromContext != uint(userId) {
		_ = c.Error(errorHandler.ErrUnauthorized)
		return
	}

	resumes, getResumeErr := rh.resumeUseCase.GetResumeByApplicant(uint(userId))
	if getResumeErr != nil {
		_ = c.Error(getResumeErr)
		return
	}

	c.JSON(http.StatusOK, resumes)
}

func (rh *ResumeHandler) CreateResume(c *gin.Context) {
	userId, getUserErr := rh.userHandler.GetUserId(c)
	if getUserErr != nil {
		_ = c.Error(getUserErr)
		return
	}

	var input models.Resume
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	creatingErr := rh.resumeUseCase.CreateResume(&input, userId)
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

	isAccessAllowed := rh.isResumeAvailable(c, uint(id))
	if isAccessAllowed != nil {
		_ = c.Error(isAccessAllowed)
		return
	}

	var input models.Resume
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	updateErr := rh.resumeUseCase.UpdateResume(uint(id), &input)
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

	isAccessAllowed := rh.isResumeAvailable(c, uint(id))
	if isAccessAllowed != nil {
		_ = c.Error(isAccessAllowed)
		return
	}

	deleteErr := rh.resumeUseCase.DeleteResume(uint(id))
	if deleteErr != nil {
		_ = c.Error(deleteErr)
		return
	}

	c.Status(http.StatusOK)
}

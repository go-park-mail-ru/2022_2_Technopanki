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

func (rh *ResumeHandler) isAccessAllowed(id uint, c *gin.Context) error {
	userId, userErr := rh.userHandler.GetUserId(c)
	if userErr != nil {
		return userErr
	}

	resume, getResumeErr := rh.resumeUseCase.Get(id)
	if getResumeErr != nil {
		return getResumeErr
	}

	if resume.UserAccountId != userId {
		return errorHandler.ErrForbidden
	}
	return nil
}

func (rh *ResumeHandler) Get(c *gin.Context) {
	id, paramErr := strconv.Atoi(c.Param("id"))
	if paramErr != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	resume, getResumeErr := rh.resumeUseCase.Get(uint(id))
	if getResumeErr != nil {
		_ = c.Error(getResumeErr)
		return
	}

	c.JSON(http.StatusOK, resume)
}

func (rh *ResumeHandler) GetByApplicant(c *gin.Context) {
	userId, paramErr := strconv.Atoi(c.Param("user_id"))
	if paramErr != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}
	resumes, getResumeErr := rh.resumeUseCase.GetByApplicant(uint(userId))
	if getResumeErr != nil {
		_ = c.Error(getResumeErr)
		return
	}

	c.JSON(http.StatusOK, resumes)
}

func (rh *ResumeHandler) Create(c *gin.Context) {
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

	resumeId, creatingErr := rh.resumeUseCase.Create(&input, userId)
	if creatingErr != nil {
		_ = c.Error(creatingErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": resumeId, "id2": input.ID})
}

func (rh *ResumeHandler) Update(c *gin.Context) {
	id, paramErr := strconv.Atoi(c.Param("id"))
	if paramErr != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	isAccessAllowed := rh.isAccessAllowed(uint(id), c)
	if isAccessAllowed != nil {
		_ = c.Error(isAccessAllowed)
		return
	}

	var input models.Resume
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	updateErr := rh.resumeUseCase.Update(uint(id), &input)
	if updateErr != nil {
		_ = c.Error(updateErr)
		return
	}

	c.Status(http.StatusOK)
}

func (rh *ResumeHandler) Delete(c *gin.Context) {
	id, paramErr := strconv.Atoi(c.Param("id"))
	if paramErr != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	isAccessAllowed := rh.isAccessAllowed(uint(id), c)
	if isAccessAllowed != nil {
		_ = c.Error(isAccessAllowed)
		return
	}

	deleteErr := rh.resumeUseCase.Delete(uint(id))
	if deleteErr != nil {
		_ = c.Error(deleteErr)
		return
	}

	c.Status(http.StatusOK)
}

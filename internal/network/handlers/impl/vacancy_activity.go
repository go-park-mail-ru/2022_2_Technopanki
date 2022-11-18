package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VacancyActivityHandler struct {
	vacancyActivityUseCase usecases.VacancyActivity
}

func NewVacancyActivityHandler(useCases *usecases.UseCases) *VacancyActivityHandler {
	return &VacancyActivityHandler{vacancyActivityUseCase: useCases.VacancyActivity}
}

func (vah *VacancyActivityHandler) GetAllVacancyApplies(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	applies, getAllErr := vah.vacancyActivityUseCase.GetAllVacancyApplies(id)
	if getAllErr != nil {
		_ = c.Error(getAllErr)
		return
	}
	c.JSON(http.StatusOK, models.GetAllAppliesResponce{
		Data: applies,
	})
}

func (vah *VacancyActivityHandler) ApplyForVacancy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	email, contextErr := getEmailFromContext(c)
	if contextErr != nil {
		_ = c.Error(contextErr)
		return
	}

	var input models.VacancyActivity
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	applyErr := vah.vacancyActivityUseCase.ApplyForVacancy(email, id, &input)
	if applyErr != nil {
		_ = c.Error(applyErr)
		return
	}

	c.Status(http.StatusOK)
}

func (vah *VacancyActivityHandler) GetAllUserApplies(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	applies, getErr := vah.vacancyActivityUseCase.GetAllUserApplies(userId)
	if getErr != nil {
		_ = c.Error(getErr)
		return
	}
	c.JSON(http.StatusOK, models.GetAllAppliesResponce{
		Data: applies,
	})
}

func (vah *VacancyActivityHandler) DeleteUserApply(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	email, contextErr := getEmailFromContext(c)
	if contextErr != nil {
		_ = c.Error(contextErr)
		return
	}

	deleteErr := vah.vacancyActivityUseCase.DeleteUserApply(email, id)
	if deleteErr != nil {
		_ = c.Error(deleteErr)
		return
	}
	c.Status(http.StatusOK)
}

package handlers

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
	userHandler            UserH
}

func newVacancyActivityHandler(useCases *usecases.UseCases, userHandler *UserHandler) *VacancyActivityHandler {
	return &VacancyActivityHandler{vacancyActivityUseCase: useCases.VacancyActivity, userHandler: userHandler}
}

type GetAllVacancyAppliesResponce struct {
	Data []*models.VacancyActivity `json:"data"`
}

type GetAllUserAppliesResponce struct {
	Data []models.VacancyActivity `json:"data"`
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
	c.JSON(http.StatusOK, GetAllVacancyAppliesResponce{
		applies,
	})
}

func (vah *VacancyActivityHandler) ApplyForVacancy(c *gin.Context) {
	userId, getUserIdErr := vah.userHandler.GetUserId(c)
	if getUserIdErr != nil {
		_ = c.Error(getUserIdErr)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}
	var input models.VacancyActivity
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	applyErr := vah.vacancyActivityUseCase.ApplyForVacancy(userId, id, &input)
	if err != nil {
		_ = c.Error(applyErr)
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
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
	c.JSON(http.StatusOK, GetAllUserAppliesResponce{
		applies,
	})
}

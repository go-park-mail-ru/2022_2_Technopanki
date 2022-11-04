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

func newVacancyActivityHandler(useCases *usecases.UseCases, handlers *Handlers) *VacancyActivityHandler {
	return &VacancyActivityHandler{vacancyActivityUseCase: useCases.VacancyActivity, userHandler: handlers.UserHandler}
}

type GetAllVacancyAppliesResponce struct {
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
	userId := vah.userHandler.GetUserId(c)
	var input models.VacancyActivity
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	err := vah.vacancyActivityUseCase.ApplyForVacancy(userId, &input)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

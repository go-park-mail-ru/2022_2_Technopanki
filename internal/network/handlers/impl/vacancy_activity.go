package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/network/handlers/utils"
	"HeadHunter/internal/usecases"
	"HeadHunter/pkg/errorHandler"
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
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

	applies, getAllErr := vah.vacancyActivityUseCase.GetAllVacancyApplies(uint(id))
	if getAllErr != nil {
		_ = c.Error(getAllErr)
		return
	}
	var appliesResponse models.GetAllAppliesResponce
	appliesResponse.Data = applies
	appliesJson, err := appliesResponse.MarshalJSON()
	if err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", appliesJson)
}

func (vah *VacancyActivityHandler) ApplyForVacancy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	email, contextErr := utils.GetEmailFromContext(c)
	if contextErr != nil {
		_ = c.Error(contextErr)
		return
	}

	var input models.VacancyActivity
	if err := easyjson.UnmarshalFromReader(c.Request.Body, &input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	applyErr := vah.vacancyActivityUseCase.ApplyForVacancy(email, uint(id), &input)
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

	applies, getErr := vah.vacancyActivityUseCase.GetAllUserApplies(uint(userId))
	if getErr != nil {
		_ = c.Error(getErr)
		return
	}
	var appliesResponse models.GetAllAppliesResponce
	appliesResponse.Data = applies
	appliesJson, err := appliesResponse.MarshalJSON()
	if err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", appliesJson)
}

func (vah *VacancyActivityHandler) DeleteUserApply(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	email, contextErr := utils.GetEmailFromContext(c)
	if contextErr != nil {
		_ = c.Error(contextErr)
		return
	}

	deleteErr := vah.vacancyActivityUseCase.DeleteUserApply(email, uint(id))
	if deleteErr != nil {
		_ = c.Error(deleteErr)
		return
	}
	c.Status(http.StatusOK)
}

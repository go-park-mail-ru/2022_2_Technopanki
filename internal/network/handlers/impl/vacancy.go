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
	"strings"
)

type VacancyHandler struct {
	vacancyUseCase usecases.Vacancy
}

func NewVacancyHandler(useCases *usecases.UseCases) *VacancyHandler {
	return &VacancyHandler{vacancyUseCase: useCases.Vacancy}
}

func (vh *VacancyHandler) GetAllVacancies(c *gin.Context) {
	var vacancies []*models.Vacancy
	var getAllErr error
	var filters models.VacancyFilter

	title := c.Query("search")
	if title != "" {
		filters.Title = title
	}

	experience := c.Query("experience")
	if experience != "" {
		filters.Experience = experience
	}

	city := c.Query("city")
	if city != "" {
		filters.Location = city
	}

	format := c.Query("format")
	if format != "" {
		filters.Format = format
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

	vacancies, getAllErr = vh.vacancyUseCase.GetAll(filters)
	if getAllErr != nil {
		_ = c.Error(getAllErr)
		return
	}
	var vacanciesResponse models.GetAllVacanciesResponcePointer
	vacanciesResponse.Data = vacancies
	vacanciesJson, err := vacanciesResponse.MarshalJSON()
	if err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", vacanciesJson)
}

func (vh *VacancyHandler) GetVacancyById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}
	vacancy, err := vh.vacancyUseCase.GetById(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}
	vacancyJson, errJson := vacancy.MarshalJSON()
	if errJson != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", vacancyJson)

}

func (vh *VacancyHandler) GetPreviewVacanciesByEmployer(c *gin.Context) {
	companyId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}
	var vacancies models.VacancyPreviews
	vacancies, GetErr := vh.vacancyUseCase.GetPreviewVacanciesByEmployer(uint(companyId))
	if GetErr != nil {
		_ = c.Error(GetErr)
		return
	}
	vacanciesJson, errJson := vacancies.MarshalJSON()
	if errJson != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", vacanciesJson)
}

func (vh *VacancyHandler) GetUserVacancies(c *gin.Context) {
	companyId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}
	vacancies, GetErr := vh.vacancyUseCase.GetByUserId(uint(companyId))
	if GetErr != nil {
		_ = c.Error(GetErr)
		return
	}
	var vacanciesResponse models.GetAllVacanciesResponcePointer
	vacanciesResponse.Data = vacancies
	vacanciesJson, err := vacanciesResponse.MarshalJSON()
	if err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", vacanciesJson)

}

func (vh *VacancyHandler) CreateVacancy(c *gin.Context) {
	email, contextErr := utils.GetEmailFromContext(c)
	if contextErr != nil {
		_ = c.Error(contextErr)
		return
	}

	var input models.Vacancy
	if err := easyjson.UnmarshalFromReader(c.Request.Body, &input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	id, err := vh.vacancyUseCase.Create(email, &input)
	if err != nil {
		_ = c.Error(err)
		return
	}
	var response models.Response
	response.ID = id
	responseJson, errJson := response.MarshalJSON()
	if errJson != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", responseJson)
}

func (vh *VacancyHandler) DeleteVacancy(c *gin.Context) {
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

	deleteErr := vh.vacancyUseCase.Delete(email, uint(id))
	if deleteErr != nil {
		_ = c.Error(deleteErr)
		return
	}
	c.Status(http.StatusOK)
}

func (vh *VacancyHandler) UpdateVacancy(c *gin.Context) {

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

	var input models.Vacancy
	if err := easyjson.UnmarshalFromReader(c.Request.Body, &input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	updateErr := vh.vacancyUseCase.Update(email, uint(id), &input)
	if updateErr != nil {
		_ = c.Error(updateErr)
		return
	}
	c.Status(http.StatusOK)
}

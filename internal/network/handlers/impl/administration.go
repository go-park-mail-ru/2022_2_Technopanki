package impl

import (
	"HeadHunter/internal/network/handlers/utils"
	"HeadHunter/internal/usecases"
	"HeadHunter/pkg/errorHandler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdministrationHandler struct {
	adminUseCase usecases.Administration
}

func NewAdministrationHandler(adminUseCase usecases.Administration) *AdministrationHandler {
	return &AdministrationHandler{adminUseCase: adminUseCase}
}

func (ah *AdministrationHandler) GetStyles(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "node_modules/bootstrap/dist/css/bootstrap.min.css")
}

func (ah *AdministrationHandler) GetForbiddenPage(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "static/html/403Forbidden.html")
}

func (ah *AdministrationHandler) GetMainPage(c *gin.Context) {
	email, contextErr := utils.GetEmailFromContext(c)
	if contextErr != nil {
		if contextErr == errorHandler.ErrUnauthorized {
			ah.GetAuthPage(c)
			return
		}
		_ = c.Error(contextErr)
		return
	}
	page, getErr := ah.adminUseCase.GetMainPage(email)
	if getErr != nil {
		_ = c.Error(getErr)
		return
	}

	c.Data(http.StatusOK, "text/html", page)
}

func (ah *AdministrationHandler) GetAuthPage(c *gin.Context) {
	//email, contextErr := utils.GetEmailFromContext(c)
	//if contextErr != nil && contextErr != errorHandler.ErrUnauthorized {
	//	_ = c.Error(contextErr)
	//	return
	//}
	//page, getErr := ah.adminUseCase.GetAuthPage(email)
	//if getErr != nil {
	//	_ = c.Error(getErr)
	//	return
	//}
	//
	//c.Data(http.StatusOK, "text/html", page)
	//http.ServeFile(c.Writer, c.Request, "./static/html/admin/auth.html")
	page, getErr := ah.adminUseCase.GetAuthPage()
	if getErr != nil {
		_ = c.Error(getErr)
		return
	}
	c.Data(http.StatusOK, "text/html", page)
}

func (ah *AdministrationHandler) GetResumesPage(c *gin.Context) {
	email, contextErr := utils.GetEmailFromContext(c)
	if contextErr != nil {
		if contextErr == errorHandler.ErrUnauthorized {
			ah.GetAuthPage(c)
			return
		}
		_ = c.Error(contextErr)
		return
	}
	page, getErr := ah.adminUseCase.GetResumesPage(email)
	if getErr != nil {
		_ = c.Error(getErr)
		return
	}

	c.Data(http.StatusOK, "text/html", page)
}

func (ah *AdministrationHandler) GetVacanciesPage(c *gin.Context) {
	email, contextErr := utils.GetEmailFromContext(c)
	if contextErr != nil {
		if contextErr == errorHandler.ErrUnauthorized {
			ah.GetAuthPage(c)
			return
		}
		_ = c.Error(contextErr)
		return
	}
	page, getErr := ah.adminUseCase.GetVacanciesPage(email)
	if getErr != nil {
		_ = c.Error(getErr)
		return
	}

	c.Data(http.StatusOK, "text/html", page)
}

func (ah *AdministrationHandler) GetApplicantsPage(c *gin.Context) {
	email, contextErr := utils.GetEmailFromContext(c)
	if contextErr != nil {
		if contextErr == errorHandler.ErrUnauthorized {
			ah.GetAuthPage(c)
			return
		}
		_ = c.Error(contextErr)
		return
	}
	page, getErr := ah.adminUseCase.GetApplicantsPage(email)
	if getErr != nil {
		_ = c.Error(getErr)
		return
	}

	c.Data(http.StatusOK, "text/html", page)
}

func (ah *AdministrationHandler) GetEmployersPage(c *gin.Context) {
	email, contextErr := utils.GetEmailFromContext(c)
	if contextErr != nil {
		if contextErr == errorHandler.ErrUnauthorized {
			ah.GetAuthPage(c)
			return
		}
		_ = c.Error(contextErr)
		return
	}
	page, getErr := ah.adminUseCase.GetEmployersPage(email)
	if getErr != nil {
		_ = c.Error(getErr)
		return
	}

	c.Data(http.StatusOK, "text/html", page)
}

package handlers

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", signUp)
		auth.POST("/sign-in", signIn)
	}
	api := router.Group("/api", UserIdentity)
	{
		vacancies := api.Group("/vacancy")
		{
			vacancies.GET("/", GetVacancies)
			vacancies.GET("/:id", GetVacancyByID)
			vacancies.POST("/", PostVacancies)
			//vacancies.PUT("/id")
			//vacancies.DELETE("/id")
		}
	}

	return router
}

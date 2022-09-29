package network

import (
	"HeadHunter/internal/network/handlers"
	"HeadHunter/internal/network/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handlers.SignUp)
		auth.POST("/sign-in", handlers.SignIn)
	}

	api := router.Group("/api", middleware.Session)
	{
		vacancies := api.Group("/vacancy")
		{
			vacancies.GET("/", handlers.GetVacancies)
			vacancies.GET("/:id", handlers.GetVacancyByID)
		}
	}

	return router
}

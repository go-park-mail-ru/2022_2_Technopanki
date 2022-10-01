package network

import (
	"HeadHunter/internal/network/handlers"
	"HeadHunter/internal/network/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handlers.SignUp)
		auth.POST("/sign-in", handlers.SignIn)
		auth.POST("/logout", handlers.Logout)
	}

	api := router.Group("/api", middleware.Session)
	{
		vacancies := api.Group("/vacancy")
		{
			vacancies.GET("/", handlers.GetVacancies, middleware.ErrorHandler())
		}
	}

	return router
}

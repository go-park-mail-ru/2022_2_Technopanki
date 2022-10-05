package network

import (
	_ "HeadHunter/docs"
	"HeadHunter/internal/network/handlers"
	"HeadHunter/internal/network/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handlers.SignUp, middleware.ErrorHandler())
		auth.POST("/sign-in", handlers.SignIn, middleware.ErrorHandler())
		auth.POST("/logout", handlers.Logout, middleware.Session, middleware.ErrorHandler())
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

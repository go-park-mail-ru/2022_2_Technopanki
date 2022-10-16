package network

import (
	_ "HeadHunter/docs"
	"HeadHunter/internal/network/handlers"
	"HeadHunter/internal/network/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(h *handlers.Handler) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(404, gin.H{"error": "invalid route (check HTTP Methods)"})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.GET("/", middleware.Session, h.AuthCheck, middleware.ErrorHandler())
		auth.POST("/sign-up", h.SignUp, middleware.ErrorHandler())
		auth.POST("/sign-in", h.SignIn, middleware.ErrorHandler())
		auth.POST("/logout", middleware.Session, h.Logout, middleware.ErrorHandler())
	}

	api := router.Group("/api")
	{
		vacancies := api.Group("/vacancy")
		{
			vacancies.GET("/", handlers.GetVacancies, middleware.ErrorHandler())
		}
	}

	return router
}

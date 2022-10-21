package network

import (
	_ "HeadHunter/docs"
	"HeadHunter/internal/network/handlers"
	"HeadHunter/internal/network/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(h *handlers.Handlers) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(404, gin.H{"error": "invalid route (check HTTP Methods)"})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.GET("/", middleware.Session, h.UserHandler.AuthCheck, middleware.ErrorHandler())
		auth.POST("/sign-up", h.UserHandler.SignUp, middleware.ErrorHandler())
		auth.POST("/sign-in", h.UserHandler.SignIn, middleware.ErrorHandler())
		auth.POST("/logout", middleware.Session, h.UserHandler.Logout, middleware.ErrorHandler())
	}

	api := router.Group("/api")
	{
		vacancies := api.Group("/vacancy")
		{
			vacancies.GET("/", handlers.GetVacancies, middleware.ErrorHandler()) //TODO заменить на строку ниже
			//vacancies.GET("/", h.VacancyHandler.Get, middleware.ErrorHandler())
			//vacancies.POST("/", h.VacancyHandler.Create, middleware.ErrorHandler())
			//vacancies.PUT("/", h.VacancyHandler.Update, middleware.ErrorHandler())
			//vacancies.DELETE("/", h.VacancyHandler.Delete, middleware.ErrorHandler())
		}
		//
		//resumes := api.Group("/resume")
		//{
		//	resumes.GET("/", h.ResumeHandler.Get, middleware.ErrorHandler())
		//	resumes.POST("/", h.ResumeHandler.Create, middleware.ErrorHandler())
		//	resumes.PUT("/", h.ResumeHandler.Update, middleware.ErrorHandler())
		//	resumes.DELETE("/", h.ResumeHandler.Delete, middleware.ErrorHandler())
		//}
	}

	return router
}

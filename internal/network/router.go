package network

import (
	_ "HeadHunter/docs"
	"HeadHunter/internal/network/handlers"
	"HeadHunter/internal/network/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(h *handlers.Handlers, sessionMW *middleware.SessionMiddleware) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(404, gin.H{"error": "invalid route (check HTTP Methods)"})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.GET("/", sessionMW.Session, h.UserHandler.AuthCheck, middleware.ErrorHandler())
		auth.POST("/sign-up", h.UserHandler.SignUp, middleware.ErrorHandler())
		auth.POST("/sign-in", h.UserHandler.SignIn, middleware.ErrorHandler())
		auth.POST("/logout", sessionMW.Session, h.UserHandler.Logout, middleware.ErrorHandler())
	}

	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.PUT("/", sessionMW.Session, h.UserHandler.UpgradeUser, middleware.ErrorHandler())
		}
		vacancies := api.Group("/vacancy")
		{
			//vacancies.GET("/", handlers.GetVacancies, middleware.ErrorHandler()) //TODO заменить на строку ниже
			vacancies.GET("/", h.VacancyHandler.GetAllVacancies, middleware.ErrorHandler())
			vacancies.GET("/:id", h.VacancyHandler.GetVacancyById, middleware.ErrorHandler())
			vacancies.GET("/company/:id", h.VacancyHandler.GetUserVacancies, middleware.ErrorHandler())
			vacancies.POST("/", sessionMW.Session, h.VacancyHandler.CreateVacancy, middleware.ErrorHandler())
			vacancies.PUT("/:id", sessionMW.Session, h.VacancyHandler.UpdateVacancy, middleware.ErrorHandler())
			vacancies.DELETE("/:id", sessionMW.Session, h.VacancyHandler.DeleteVacancy, middleware.ErrorHandler())
			vacancies.POST("/apply", sessionMW.Session, h.VacancyActivityHandler.ApplyForVacancy, middleware.ErrorHandler())
			vacancies.GET("/applies/:id", sessionMW.Session, h.VacancyActivityHandler.GetAllVacancyApplies, middleware.ErrorHandler())
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

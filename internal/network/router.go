package network

import (
	"HeadHunter/configs"
	_ "HeadHunter/docs"
	"HeadHunter/internal/network/handlers"
	"HeadHunter/internal/network/middleware"
	"HeadHunter/internal/network/ws"
	"HeadHunter/pkg/errorHandler"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(h *handlers.Handlers, sessionMW *middleware.SessionMiddleware, cfg *configs.Config, wsPool *ws.Pool) *gin.Engine {
	router := gin.Default()

	monitor := ginmetrics.GetMonitor()
	monitor.SetMetricPath(cfg.MetricPath)
	monitor.SetSlowTime(10)
	monitor.Use(router)

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(404, gin.H{"error": "invalid route (check HTTP Methods)"})
	})

	router.Use(middleware.CORSMiddleware())

	initCSRF(router, cfg.Security)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	wsMiddleware := middleware.NewWebsocketMiddleware(wsPool)

	auth := router.Group("/auth")
	{
		auth.GET("/", sessionMW.Session, h.UserHandler.AuthCheck, errorHandler.Middleware())
		auth.POST("/sign-up", h.UserHandler.SignUp, errorHandler.Middleware())
		auth.POST("/sign-in", h.UserHandler.SignIn, errorHandler.Middleware())
		auth.POST("/logout", sessionMW.Session, h.UserHandler.Logout, errorHandler.Middleware())
		auth.POST("/confirm", h.UserHandler.ConfirmUser, errorHandler.Middleware())
	}

	mail := router.Group("/mail")
	{
		mail.POST("/code/:email", h.MailHandler.SendConfirmCode, errorHandler.Middleware())
	}

	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.GET("/:id", sessionMW.Session, h.UserHandler.GetUser, errorHandler.Middleware())
			user.GET("/employers", h.UserHandler.GetAllEmployers, errorHandler.Middleware())
			user.GET("/applicants", h.UserHandler.GetAllApplicants, errorHandler.Middleware())
			user.GET("/safety/:id", h.UserHandler.GetUserSafety, errorHandler.Middleware())
			user.GET("/preview/:id", h.UserHandler.GetPreview, errorHandler.Middleware())
			user.POST("/", sessionMW.Session, h.UserHandler.UpdateUser, errorHandler.Middleware())
			user.POST("/password", sessionMW.Session, h.UserHandler.UpdatePassword, errorHandler.Middleware())
			user.GET("/mailing", sessionMW.Session, h.UserHandler.GetMailing, errorHandler.Middleware())
			image := user.Group("/image")
			{
				image.POST("/", sessionMW.Session, h.UserHandler.UploadUserImage, errorHandler.Middleware())
				image.DELETE("/", sessionMW.Session, h.UserHandler.DeleteUserImage, errorHandler.Middleware())
			}
		}

		vacancies := api.Group("/vacancy")
		{
			vacancies.GET("", h.VacancyHandler.GetAllVacancies, errorHandler.Middleware())
			vacancies.GET("/:id", h.VacancyHandler.GetVacancyById, errorHandler.Middleware())
			vacancies.GET("/employer/preview/:id", h.VacancyHandler.GetPreviewVacanciesByEmployer, errorHandler.Middleware())
			vacancies.GET("/company/:id", h.VacancyHandler.GetUserVacancies, errorHandler.Middleware())
			vacancies.POST("/", sessionMW.Session, h.VacancyHandler.CreateVacancy, errorHandler.Middleware())
			vacancies.PUT("/:id", sessionMW.Session, h.VacancyHandler.UpdateVacancy, errorHandler.Middleware())
			vacancies.DELETE("/:id", sessionMW.Session, h.VacancyHandler.DeleteVacancy, errorHandler.Middleware())
			vacancies.POST("/apply/:id", sessionMW.Session, h.VacancyActivityHandler.ApplyForVacancy, wsMiddleware.Send, errorHandler.Middleware())
			vacancies.GET("/applies/:id", h.VacancyActivityHandler.GetAllVacancyApplies, errorHandler.Middleware())
			vacancies.GET("/user_applies/:id", h.VacancyActivityHandler.GetAllUserApplies, errorHandler.Middleware())
			vacancies.POST("/favorites/:id", sessionMW.Session, h.VacancyHandler.AddVacancyToFavorites, errorHandler.Middleware())
			vacancies.GET("/favorites", sessionMW.Session, h.VacancyHandler.GetUserFavoriteVacancies, errorHandler.Middleware())
			vacancies.DELETE("/favorites/:id", sessionMW.Session, h.VacancyHandler.DeleteVacancyFromFavorites, errorHandler.Middleware())
			vacancies.GET("/favorites/check/:id", sessionMW.Session, h.VacancyHandler.CheckFavoriteVacancy, errorHandler.Middleware())

		}

		resumes := api.Group("/resume")
		{
			resumes.GET("", h.ResumeHandler.GetAllResumes, errorHandler.Middleware())
			resumes.GET("/:id", h.ResumeHandler.GetResume, errorHandler.Middleware())
			resumes.GET("/applicant/:user_id", h.ResumeHandler.GetResumeByApplicant, errorHandler.Middleware())
			resumes.GET("/applicant/preview/:user_id", h.ResumeHandler.GetPreviewResumeByApplicant, errorHandler.Middleware())
			resumes.GET("/pdf/:id", sessionMW.Session, h.ResumeHandler.GetResumeInPDF, wsMiddleware.Send, errorHandler.Middleware())
			resumes.POST("/", sessionMW.Session, h.ResumeHandler.CreateResume, errorHandler.Middleware())
			resumes.PUT("/:id", sessionMW.Session, h.ResumeHandler.UpdateResume, errorHandler.Middleware())
			resumes.DELETE("/:id", sessionMW.Session, h.ResumeHandler.DeleteResume, errorHandler.Middleware())
		}

		notifications := api.Group("/notification")
		{
			notifications.GET("/ws", sessionMW.Session, wsPool.Connect, errorHandler.Middleware())
			notifications.GET("/user", sessionMW.Session, h.NotificationHandler.GetNotifications, errorHandler.Middleware())
			notifications.PUT("/read", sessionMW.Session, h.NotificationHandler.ReadAllNotifications, errorHandler.Middleware())
			notifications.PUT("/read/:id", sessionMW.Session, h.NotificationHandler.ReadNotification, errorHandler.Middleware())
			notifications.DELETE("/clear", sessionMW.Session, h.NotificationHandler.ClearNotifications, errorHandler.Middleware())
		}
	}

	return router
}

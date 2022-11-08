package network

import (
	"HeadHunter/configs"
	_ "HeadHunter/docs"
	"HeadHunter/internal/network/handlers"
	"HeadHunter/internal/network/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(h *handlers.Handlers, sessionMW *middleware.SessionMiddleware, cfg *configs.Config) *gin.Engine {
	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(404, gin.H{"error": "invalid route (check HTTP Methods)"})
	})

	router.Use(middleware.CORSMiddleware())

	//store := cookie.NewStore([]byte("secret"))
	//router.Use(sessions.Sessions("mySession", store))
	//router.Use(csrf.Middleware(csrf.Options{
	//	Secret: cfg.Security.Secret,
	//	ErrorFunc: func(c *gin.Context) {
	//		_ = c.Error(errorHandler.CSRFTokenMismatch)
	//		return
	//	},
	//}))

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
			user.GET("/:id", sessionMW.Session, h.UserHandler.GetUser, middleware.ErrorHandler())
			user.GET("/safety/:id", h.UserHandler.GetUserSafety, middleware.ErrorHandler())
			user.GET("/preview/:id", h.UserHandler.GetPreview, middleware.ErrorHandler())
			user.POST("/", sessionMW.Session, h.UserHandler.UpdateUser, middleware.ErrorHandler())
			image := user.Group("/image")
			{
				image.POST("/", sessionMW.Session, h.UserHandler.UploadUserImage, middleware.ErrorHandler())
				image.DELETE("/", sessionMW.Session, h.UserHandler.DeleteUserImage, middleware.ErrorHandler())
			}
		}
		vacancies := api.Group("/vacancy")
		{
			vacancies.GET("/", handlers.GetVacancies, middleware.ErrorHandler()) //TODO заменить на строку ниже
			//vacancies.GET("/", h.VacancyHandler.GetResume, middleware.ErrorHandler())
			//vacancies.POST("/", h.VacancyHandler.CreateResume, middleware.ErrorHandler())
			//vacancies.PUT("/", h.VacancyHandler.UpdateResume, middleware.ErrorHandler())
			//vacancies.DELETE("/", h.VacancyHandler.DeleteResume, middleware.ErrorHandler())
		}
		//
		resumes := api.Group("/resume")
		{
			resumes.GET("/:id", sessionMW.Session, h.ResumeHandler.GetResume, middleware.ErrorHandler())
			resumes.GET("/applicant/:user_id", sessionMW.Session, h.ResumeHandler.GetResumeByApplicant, middleware.ErrorHandler())
			resumes.POST("/", sessionMW.Session, h.ResumeHandler.CreateResume, middleware.ErrorHandler())
			resumes.PUT("/:id", sessionMW.Session, h.ResumeHandler.UpdateResume, middleware.ErrorHandler())
			resumes.DELETE("/:id", sessionMW.Session, h.ResumeHandler.DeleteResume, middleware.ErrorHandler())
		}
	}

	return router
}

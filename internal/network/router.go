package network

import (
	"HeadHunter/configs"
	_ "HeadHunter/docs"
	"HeadHunter/internal/errorHandler"
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

	protected := router.Group("/protected")
	{
		protected.GET("/", func(c *gin.Context) {
			token, err := c.GetRawData()
			if err != nil {
				_ = c.Error(errorHandler.ErrBadRequest)
				return
			}
			c.SetCookie("X-CSRF-Token", string(token), 0, "/",
				cfg.Domain, cfg.Cookie.Secure, cfg.Cookie.HTTPOnly)
		}, middleware.ErrorHandler())

		protected.POST("/", func(c *gin.Context) {
			c.String(200, "CSRF token is valid")
		})
	}

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
			vacancies.GET("/", h.VacancyHandler.GetAllVacancies, middleware.ErrorHandler())
			vacancies.GET("/:id", h.VacancyHandler.GetVacancyById, middleware.ErrorHandler())
			vacancies.GET("/company/:id", h.VacancyHandler.GetUserVacancies, middleware.ErrorHandler())
			vacancies.POST("/", sessionMW.Session, h.VacancyHandler.CreateVacancy, middleware.ErrorHandler())
			vacancies.PUT("/:id", sessionMW.Session, h.VacancyHandler.UpdateVacancy, middleware.ErrorHandler())
			vacancies.DELETE("/:id", sessionMW.Session, h.VacancyHandler.DeleteVacancy, middleware.ErrorHandler())
			vacancies.POST("/apply", sessionMW.Session, h.VacancyActivityHandler.ApplyForVacancy, middleware.ErrorHandler())
			vacancies.GET("/applies/:id", h.VacancyActivityHandler.GetAllVacancyApplies, middleware.ErrorHandler())
			vacancies.GET("/user_applies/:id", h.VacancyActivityHandler.GetAllUserApplies, middleware.ErrorHandler())

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

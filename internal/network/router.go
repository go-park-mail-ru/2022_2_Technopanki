package network

import (
	"HeadHunter/configs"
	_ "HeadHunter/docs"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/network/handlers"
	"HeadHunter/internal/network/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	csrf "github.com/utrack/gin-csrf"
)

func InitRoutes(h *handlers.Handlers, sessionMW *middleware.SessionMiddleware, cfg *configs.Config) *gin.Engine {
	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(404, gin.H{"error": "invalid route (check HTTP Methods)"})
	})

	router.Use(middleware.CORSMiddleware())

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mySession", store))
	router.Use(csrf.Middleware(csrf.Options{
		Secret: cfg.Security.Secret,
		ErrorFunc: func(c *gin.Context) {
			_ = c.Error(errorHandler.CSRFTokenMismatch)
			return
		},
	}))

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

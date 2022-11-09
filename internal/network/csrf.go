package network

import (
	"HeadHunter/configs"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func initCSRF(router *gin.Engine, cfg configs.SecurityConfig) {
	if !cfg.CsrfMode {
		return
	}
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("csrfSession", store))
	router.Use(csrf.Middleware(csrf.Options{
		Secret: cfg.Secret,
		ErrorFunc: func(c *gin.Context) {
			c.String(403, "CSRF token mismatch")
			c.Abort()
		},
	}))

	protected := router.Group("/protected")
	{
		protected.GET("/", func(c *gin.Context) {
			token := csrf.GetToken(c)
			c.Request.Header.Add("X-CSRF-Token", token)
			c.JSON(200, gin.H{"token": csrf.GetToken(c)})
		})

		protected.POST("/", func(c *gin.Context) {
			c.String(200, "CSRF token is valid")
		})
	}
}

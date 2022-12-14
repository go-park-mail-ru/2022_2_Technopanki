package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"http://localhost:8000", "http://95.163.208.72:8000",
		"http://localhost:80", "http://95.163.208.72:80"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "X-CSRF-Token"}
	return cors.New(config)
}

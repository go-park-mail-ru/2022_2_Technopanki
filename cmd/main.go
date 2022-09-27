package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"main.go/internal/transport/middleware"
	"main.go/internal/transport/rest/handlers"
	"net/http"
	"os"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

func main() {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware)
	router.POST("/signup", handlers.Signup)
	router.GET("/auth", middleware.JWTMiddleware, handlers.Auth)

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "4000"
	}

	router.Run("localhost:" + port)
}

package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func newHandler(db *gorm.DB) *Handler {
	return &Handler{db}
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Employer{})

	handler := newHandler(db)

	r := gin.New()

	r.POST("/login", loginHandler)

	protected := r.Group("/", authorizationMiddleware)

	protected.GET("/employers", handler.listEmployersHandler)
	protected.POST("/employers", handler.createEmployerHandler)
	protected.DELETE("/employers/:id", handler.deleteEmployerHandler)

	r.Run()
}

type Employer struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Applicant struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Vacancy struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Salary      string `json:"salary"`
	Employer_ID string `json:"employer_id"`
	Date        string `json:"date"`
}

func loginHandler(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
	})

	ss, err := token.SignedString([]byte("MySignature"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": ss,
	})
}

func authorizationMiddleware(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(s, "Bearer ")

	if err := validateToken(token); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func (h *Handler) listEmployersHandler(c *gin.Context) {

	var employers []Employer

	if result := h.db.Find(&employers); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &employers)
}

func validateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("MySignature"), nil
	})

	return err
}

func (h *Handler) createEmployerHandler(c *gin.Context) {

	var employer Employer

	if err := c.BindJSON(&employer); err != nil {
		return
	}

	if result := h.db.Create(&employer); result.Error != nil {
		return
	}

	c.JSON(http.StatusCreated, &employer)
}

func (h *Handler) deleteEmployerHandler(c *gin.Context) {

	id := c.Param("id")

	if result := h.db.Delete(&Employer{}, id); result.Error != nil {
		return
	}

	c.Status(http.StatusNoContent)
}

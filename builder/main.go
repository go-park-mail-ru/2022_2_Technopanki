package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const port = ":22222"

func main() {
	router := gin.New()
	router.NoRoute(func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}
	
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	err := execScript()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Shutdown Server ...")
	if err = srv.Shutdown(context.Background()); err != nil {
		log.Fatalln("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}

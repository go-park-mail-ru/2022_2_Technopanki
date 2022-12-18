package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/exec"
	"time"
)

const port = ":22222"

func execScript() error {
	cmd := exec.Command("sh", "/backend/builder/buildScript.sh")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	fmt.Printf("%s\n", out.String())
	return err
}

func main() {
	err := execScript()
	if err != nil {
		log.Fatalln("exec script: ", err)
	}

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

	time.Sleep(30 * time.Minute)
	log.Println("Shutdown Server ...")
	if err = srv.Shutdown(context.Background()); err != nil {
		log.Fatalln("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}

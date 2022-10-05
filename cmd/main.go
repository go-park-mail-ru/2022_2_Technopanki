package main

import (
	"HeadHunter/internal/network"
	"log"
)

// @title Jobflow API
// @version 1.0
// @description Swagger API for Golang Project Jobflow.

// @host      localhost:8080
// @BasePath  /

func main() {
	router := network.InitRoutes()
	runErr := router.Run("localhost:8080")
	if runErr != nil {
		log.Fatal(runErr)
	}
}

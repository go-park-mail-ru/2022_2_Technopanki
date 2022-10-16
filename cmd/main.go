package main

import (
	"HeadHunter/internal/config"
	"HeadHunter/internal/network"
	"log"
)

// @title Jobflow API
// @version 1.0
// @description Swagger API for Golang Project Jobflow.

// @host      95.163.208.72:8080
// @BasePath  /

func main() {
	config.Connect()
	router := network.InitRoutes()
	runErr := router.Run(":8080")
	if runErr != nil {
		log.Fatal(runErr)
	}
}

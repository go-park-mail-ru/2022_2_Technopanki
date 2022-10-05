package main

import (
	"HeadHunter/internal/network"
	"log"
)

func main() {
	router := network.InitRoutes()
	runErr := router.Run(":8080")
	if runErr != nil {
		log.Fatal(runErr)
	}
}

package main

import (
	"HeadHunter/internal/network"
	"log"
)

func main() {
	router := network.InitRoutes()
	RunErr := router.Run("localhost:8080")
	if RunErr != nil {
		log.Fatal(RunErr)
	}
}

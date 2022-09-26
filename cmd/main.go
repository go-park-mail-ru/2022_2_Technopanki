package main

import (
	jobflow "HeadHunter/handlers"
	"log"
)

func main() {
	router := jobflow.InitRoutes()
	RunErr := router.Run("localhost:8080")
	if RunErr != nil {
		log.Fatal(RunErr)
	}
}

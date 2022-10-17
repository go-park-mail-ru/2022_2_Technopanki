package main

import (
	"HeadHunter/internal/network"
	"HeadHunter/internal/network/handlers"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/storage"
	"HeadHunter/internal/usecases"
	repositorypkg "HeadHunter/pkg/repository"
	"github.com/spf13/viper"
	"log"
)

// @title Jobflow API
// @version 1.0
// @description Swagger API for Golang Project Jobflow.

// @host      95.163.208.72:8080
// @BasePath  /
func main() {
	err := repositorypkg.Connect(repositorypkg.DBConfig{
		Host:     "jfPostgres",
		Port:     "5432",
		Username: "jobflowAdmin",
		Password: "12345",
		DBName:   "jobflowDB",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatal(err)
	}
	if configErr := initConfig(); configErr != nil {
		log.Fatal(configErr.Error())
	}

	useCase := usecases.NewUseCases(&repository.Repository{UserRepository: &storage.UserStorage})
	handler := handlers.NewHandler(useCase)
	router := network.InitRoutes(handler)
	runErr := router.Run(viper.GetString("port"))
	if runErr != nil {
		log.Fatal(runErr)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

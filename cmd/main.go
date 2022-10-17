package main

import (
	"HeadHunter/configs"
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
	var mainConfig configs.Config
	if configErr := configs.InitConfig(&mainConfig); configErr != nil {
		log.Fatal(configErr.Error())
	}

	err := repositorypkg.Connect(mainConfig.DB)
	if err != nil {
		log.Fatal(err)
	}

	useCase := usecases.NewUseCases(&repository.Repository{
		UserRepository: &storage.UserStorage,
		Cfg:            &mainConfig,
	})
	handler := handlers.NewHandlers(useCase)
	router := network.InitRoutes(handler)
	runErr := router.Run(viper.GetString("port"))
	if runErr != nil {
		log.Fatal(runErr)
	}
}

package main

import (
	"HeadHunter/configs"
	"HeadHunter/internal/network"
	"HeadHunter/internal/network/handlers"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/repository/session"
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

	_, err := repositorypkg.Connect(mainConfig.DB) //TODO добавить базу данных
	if err != nil {
		log.Fatal(err)
	}
	sessions := session.NewSessionsStore(mainConfig)
	useCase := usecases.NewUseCases(&repository.Repository{
		UserRepository: &storage.UserStorage}, //TODO добавить нормальнуб бд
		sessions, //TODO добавить нормальные сессии(Redis)
	)
	handler := handlers.NewHandlers(useCase, &mainConfig, sessions)
	router := network.InitRoutes(handler)
	runErr := router.Run(viper.GetString("port"))
	if runErr != nil {
		log.Fatal(runErr)
	}
}

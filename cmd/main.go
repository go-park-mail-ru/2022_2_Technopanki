package main

import (
	"HeadHunter/configs"
	"HeadHunter/internal/network"
	"HeadHunter/internal/network/handlers"
	"HeadHunter/internal/network/middleware"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/repository/session"
	"HeadHunter/internal/storage"
	"HeadHunter/internal/usecases"
	repositorypkg "HeadHunter/pkg/repository"
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
	client, redisErr := repositorypkg.RedisConnect(mainConfig.Redis)
	if redisErr != nil {
		log.Fatal(redisErr)
	}
	redisRepository := session.NewRedisStore(mainConfig, client)
	sessionMiddleware := middleware.NewSessionMiddleware(redisRepository)
	_, DBErr := repositorypkg.DBConnect(mainConfig.DB) //TODO добавить базу данных
	if DBErr != nil {
		log.Fatal(DBErr)
	}

	useCase := usecases.NewUseCases(&repository.Repository{
		UserRepository: &storage.UserStorage}, //TODO добавить нормальнуб бд
		redisRepository,
		&mainConfig,
	)

	handler := handlers.NewHandlers(useCase, &mainConfig, redisRepository)

	router := network.InitRoutes(handler, sessionMiddleware)
	runErr := router.Run(mainConfig.Port)
	if runErr != nil {
		log.Fatal(runErr)
	}
}

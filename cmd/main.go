package main

import (
	"HeadHunter/configs"
	"HeadHunter/internal/network"
	"HeadHunter/internal/network/handlers"
	"HeadHunter/internal/network/middleware"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/repository/session"
	"HeadHunter/internal/usecases"
	repositorypkg "HeadHunter/pkg/repository"
	"github.com/sirupsen/logrus"
)

// @title Jobflow API
// @version 1.0
// @description Swagger API for Golang Project Jobflow.

// @host      95.163.208.72:8080
// @BasePath  /
func main() {
	var mainConfig configs.Config
	if configErr := configs.InitConfig(&mainConfig); configErr != nil {
		logrus.Fatal(configErr)
	}
	redisClient, redisErr := repositorypkg.RedisConnect(&mainConfig.Redis)
	if redisErr != nil {
		logrus.Fatal(redisErr)
	}

	redisRepository := session.NewRedisStore(&mainConfig, redisClient)
	sessionMiddleware := middleware.NewSessionMiddleware(redisRepository)
	db, DBErr := repositorypkg.DBConnect(&mainConfig.DB)
	if DBErr != nil {
		logrus.Fatal(DBErr)
	}

	postgresRepository := repository.NewPostgresRepository(db)

	useCase := usecases.NewUseCases(
		postgresRepository,
		redisRepository,
		&mainConfig,
	)

	handler := handlers.NewHandlers(useCase, &mainConfig)

	router := network.InitRoutes(handler, sessionMiddleware, &mainConfig)
	runErr := router.Run(mainConfig.Port)
	if runErr != nil {
		logrus.Fatal(runErr)
	}
}

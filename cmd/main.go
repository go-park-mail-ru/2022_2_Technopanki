package main

import (
	"HeadHunter/configs"
	"HeadHunter/internal/cron"
	"HeadHunter/internal/network"
	"HeadHunter/internal/network/handlers"
	"HeadHunter/internal/network/middleware"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/repository/session"
	"HeadHunter/internal/usecases"
	"HeadHunter/internal/usecases/sender"
	repositorypkg "HeadHunter/pkg/repository"
	"fmt"
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

	senderService, senderErr := sender.NewSender(&mainConfig)

	if senderErr != nil {
		logrus.Fatal(senderErr)
	}

	useCase := usecases.NewUseCases(
		postgresRepository,
		redisRepository,
		senderService,
		&mainConfig,
	)

	handler := handlers.NewHandlers(useCase, &mainConfig)

	quit := make(chan struct{})
	go cron.ClearDBFromUnconfirmedUser(db, quit)

	router := network.InitRoutes(handler, sessionMiddleware, &mainConfig)
	runErr := router.Run(mainConfig.Port)
	if runErr != nil {
		logrus.Fatal(runErr)
	}
	fmt.Println("stop cron")
	quit <- struct{}{}
}

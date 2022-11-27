package main

import (
	"HeadHunter/auth_microservice/config"
	"HeadHunter/auth_microservice/handler"
	handlerImpl "HeadHunter/auth_microservice/handler/impl"
	repository "HeadHunter/auth_microservice/repository/impl"
	usecase "HeadHunter/auth_microservice/usecase/impl"
	repositorypkg "HeadHunter/pkg/repository"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	var sessionConfig config.SessionConfig
	if configErr := config.InitConfig(&sessionConfig); configErr != nil {
		logrus.Fatal(configErr)
	}
	redisClient, redisErr := repositorypkg.RedisConnect(&sessionConfig.Redis)
	if redisErr != nil {
		logrus.Fatal(redisErr)
	}

	redisRepository := repository.NewRedisStore(&sessionConfig, redisClient)

	sessionUseCase := usecase.NewSessionUseCase(*redisRepository)

	sessionHandler := handlerImpl.NewSessionHandler(*sessionUseCase)

	grpcSrv := grpc.NewServer()
	handler.RegisterAuthCheckerServer(grpcSrv, sessionHandler)

	listener, listenErr := net.Listen("tcp", ":8081")
	if listenErr != nil {
		log.Fatal("cant listen port: ", listenErr)
	}

	if serveErr := grpcSrv.Serve(listener); serveErr != nil {
		log.Fatal(serveErr)
	}
}

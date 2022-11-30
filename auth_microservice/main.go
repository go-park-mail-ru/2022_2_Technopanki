package main

import (
	"HeadHunter/auth_microservice/configs"
	proto "HeadHunter/auth_microservice/handler"
	handler "HeadHunter/auth_microservice/handler/impl"
	repository "HeadHunter/auth_microservice/repository/impl"
	usecase "HeadHunter/auth_microservice/usecase/impl"
	repositorypkg "HeadHunter/pkg/repository"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	var sessionConfig configs.SessionConfig
	if configErr := configs.InitConfig(&sessionConfig); configErr != nil {
		logrus.Fatal(configErr)
	}
	redisClient, redisErr := repositorypkg.RedisConnect(&sessionConfig.Redis)
	if redisErr != nil {
		logrus.Fatal(redisErr)
	}

	redisRepository := repository.NewRedisStore(&sessionConfig, redisClient)

	sessionUseCase := usecase.NewSessionUseCase(*redisRepository)

	sessionHandler := handler.NewSessionHandler(*sessionUseCase)

	grpcSrv := grpc.NewServer()
	proto.RegisterAuthCheckerServer(grpcSrv, sessionHandler)

	listener, listenErr := net.Listen("tcp", sessionConfig.Port)
	if listenErr != nil {
		log.Fatal("cant listen port: ", listenErr)
	}

	if serveErr := grpcSrv.Serve(listener); serveErr != nil {
		log.Fatal(serveErr)
	}
}

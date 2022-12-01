package main

import (
	"HeadHunter/auth_microservice/configs"
	proto "HeadHunter/auth_microservice/handler"
	handler "HeadHunter/auth_microservice/handler/impl"
	repository "HeadHunter/auth_microservice/repository/impl"
	usecase "HeadHunter/auth_microservice/usecase/impl"
	"HeadHunter/metrics"
	repositorypkg "HeadHunter/pkg/repository"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
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

	prometheus.MustRegister(metrics.SessionRequest)
	prometheus.MustRegister(metrics.SessionRequestDuration)
	http.Handle(sessionConfig.MetricPath, promhttp.Handler())
	go func() {
		log.Fatal(http.ListenAndServe(sessionConfig.MetricPort, nil))
	}()

	listener, listenErr := net.Listen("tcp", sessionConfig.Port)
	if listenErr != nil {
		log.Fatal("cant listen port: ", listenErr)
	}

	if serveErr := grpcSrv.Serve(listener); serveErr != nil {
		log.Fatal(serveErr)
	}
}

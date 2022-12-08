package main

import (
	auth_handler "HeadHunter/auth_microservice/handler"
	"HeadHunter/common/session"
	"HeadHunter/configs"
	"HeadHunter/internal/cron"
	"HeadHunter/internal/network"
	"HeadHunter/internal/network/handlers"
	"HeadHunter/internal/network/middleware"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/usecases"
	"HeadHunter/internal/usecases/mail"
	mail_handler "HeadHunter/mail_microservice/handler"
	repositorypkg "HeadHunter/pkg/repository"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
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

	grpcSession, sessionErr := grpc.Dial(
		strings.Join([]string{mainConfig.AuthMsHost, mainConfig.AuthMsPort}, ""),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if sessionErr != nil {
		logrus.Fatal(sessionErr)
	}

	sessionClient := auth_handler.NewAuthCheckerClient(grpcSession)

	sessionRepository := session.NewSessionMicroservice(sessionClient)
	sessionMiddleware := middleware.NewSessionMiddleware(sessionRepository)
	db, DBErr := repositorypkg.DBConnect(&mainConfig.DB)
	if DBErr != nil {
		logrus.Fatal(DBErr)
	}

	postgresRepository := repository.NewPostgresRepository(db)

	grpcMail, mailErr := grpc.Dial(
		strings.Join([]string{mainConfig.MailMsHost, mainConfig.MailMsPort}, ""),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if mailErr != nil {
		logrus.Fatal(mailErr)
	}

	mailClient := mail_handler.NewMailServiceClient(grpcMail)

	mailService := mail.NewMailService(mailClient)

	useCase := usecases.NewUseCases(
		postgresRepository,
		sessionRepository,
		&mainConfig,
		mailService,
	)

	handler := handlers.NewHandlers(useCase, &mainConfig)

	go cron.ClearDBFromUnconfirmedUser(db, &mainConfig)

	router := network.InitRoutes(handler, sessionMiddleware, &mainConfig)
	
	monitor := ginmetrics.GetMonitor()
	monitor.SetMetricPath(mainConfig.MetricPath)
	monitor.SetSlowTime(10)
	monitor.Use(router)
	runErr := router.Run(mainConfig.Port)
	if runErr != nil {
		logrus.Fatal(runErr)
	}
}

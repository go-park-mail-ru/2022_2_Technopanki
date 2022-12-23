package main

import (
	auth_handler "HeadHunter/auth_microservice/handler"
	"HeadHunter/common/session"
	"HeadHunter/configs"
	myCrons "HeadHunter/internal/cron"
	"HeadHunter/internal/network"
	"HeadHunter/internal/network/handlers"
	"HeadHunter/internal/network/middleware"
	"HeadHunter/internal/network/ws"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/usecases"
	"HeadHunter/internal/usecases/mail"
	mail_handler "HeadHunter/mail_microservice/handler"
	repositorypkg "HeadHunter/pkg/repository"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strings"
	"time"
)

// @title Jobflow API
// @version 1.0
// @description Swagger API for Golang Project Jobflow.

// @host      95.163.208.72:8080
// @BasePath  /
func main() {
	time.Sleep(5 * time.Second)
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

	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc("0 0 9 * * 0", myCrons.Mailing(postgresRepository.UserRepository, mailService))
	if err != nil {
		log.Fatal(err)
	}

	c.Start()
	defer c.Stop()

	useCase := usecases.NewUseCases(
		postgresRepository,
		sessionRepository,
		&mainConfig,
		mailService,
	)

	wsPool := ws.NewWSPool(useCase.User)

	handler := handlers.NewHandlers(useCase, &mainConfig)

	go myCrons.ClearDBFromUnconfirmedUser(db, &mainConfig)

	router := network.InitRoutes(handler, sessionMiddleware, &mainConfig, wsPool)

	runErr := router.Run(mainConfig.Port)
	if runErr != nil {
		logrus.Fatal(runErr)
	}
}

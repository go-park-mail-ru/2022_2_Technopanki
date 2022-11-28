package main

import (
	auth_handler "HeadHunter/auth_microservice/handler"
	"HeadHunter/mail_microservice/configs"
	"HeadHunter/mail_microservice/handler"
	mail_handler "HeadHunter/mail_microservice/handler/impl"
	"HeadHunter/mail_microservice/repository/session"
	usecase "HeadHunter/mail_microservice/usecase/impl"
	"HeadHunter/mail_microservice/usecase/sender"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"strings"
)

func main() {
	var mailConfig configs.Config
	configErr := configs.InitConfig(&mailConfig)
	if configErr != nil {
		logrus.Fatal(configErr)
	}
	grpcSession, sessionErr := grpc.Dial(
		strings.Join([]string{mailConfig.AuthDomain, mailConfig.AuthPort}, ""),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if sessionErr != nil {
		logrus.Fatal(sessionErr)
	}

	sessionClient := auth_handler.NewAuthCheckerClient(grpcSession)

	sessionRep := session.NewSessionMicroservice(sessionClient)

	senderService, senderErr := sender.NewSender(&mailConfig)
	if senderErr != nil {
		logrus.Fatal(senderErr)
	}

	mailService := usecase.NewMailService(sessionRep, senderService)

	mailHandler := mail_handler.NewMailHandler(mailService)

	grpcSrv := grpc.NewServer()
	handler.RegisterMailServiceServer(grpcSrv, mailHandler)

	listener, listenErr := net.Listen("tcp", mailConfig.Port)
	if listenErr != nil {
		log.Fatal("cant listen port: ", listenErr)
	}

	if serveErr := grpcSrv.Serve(listener); serveErr != nil {
		log.Fatal(serveErr)
	}
}

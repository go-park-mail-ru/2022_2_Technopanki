package main

import (
	auth_handler "HeadHunter/auth_microservice/handler"
	"HeadHunter/common/session"
	"HeadHunter/mail_microservice/configs"
	mailingCron "HeadHunter/mail_microservice/cron"
	"HeadHunter/mail_microservice/handler"
	mail_handler "HeadHunter/mail_microservice/handler/impl"
	usecase "HeadHunter/mail_microservice/usecase/impl"
	"HeadHunter/mail_microservice/usecase/sender"
	"HeadHunter/metrics"
	repositorypkg "HeadHunter/pkg/repository"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"strings"
)

func main() {
	var mailConfig configs.Config
	configErr := configs.InitConfig(&mailConfig)
	if configErr != nil {
		logrus.Fatal(configErr)
	}
	grpcSession, sessionErr := grpc.Dial(
		strings.Join([]string{mailConfig.AuthMsHost, mailConfig.AuthMsPort}, ""),
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

	db, DBErr := repositorypkg.DBConnect(&mailConfig.DB)
	if DBErr != nil {
		logrus.Fatal(DBErr)
	}
	c := cron.New(cron.WithSeconds())

	_, err := c.AddFunc("@every 20s", mailingCron.Mailing(db, mailService))
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.AddFunc("0 0 9 * * 0", mailingCron.Mailing(db, mailService))
	if err != nil {
		log.Fatal(err)
	}

	c.Start()
	defer c.Stop()

	grpcSrv := grpc.NewServer()
	handler.RegisterMailServiceServer(grpcSrv, mailHandler)

	prometheus.MustRegister(metrics.MailRequest)
	prometheus.MustRegister(metrics.MailRequestDuration)
	http.Handle(mailConfig.MetricPath, promhttp.Handler())
	go func() {
		log.Fatal(http.ListenAndServe(mailConfig.MetricPort, nil))
	}()

	listener, listenErr := net.Listen("tcp", mailConfig.Port)
	if listenErr != nil {
		log.Fatal("cant listen port: ", listenErr)
	}

	if serveErr := grpcSrv.Serve(listener); serveErr != nil {
		log.Fatal(serveErr)
	}
}

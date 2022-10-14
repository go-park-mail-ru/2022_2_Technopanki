package main

import (
	"HeadHunter/internal/network"
	"HeadHunter/internal/repository"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
)

// @title Jobflow API
// @version 1.0
// @description Swagger API for Golang Project Jobflow.

// @host      95.163.208.72:8080
// @BasePath  /
type Database struct {
	db *gorm.DB
}

func main() {
	if configErr := initConfig(); configErr != nil {
		log.Fatal(configErr.Error())
	}
	dataBase := Database{}
	db, dbErr := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "jobflowAdmin",
		Password: "12345",
		DBName:   "jobflowDB",
		SSLMode:  "disable",
	})
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	dataBase.db = db
	router := network.InitRoutes()
	runErr := router.Run(viper.GetString("port"))
	if runErr != nil {
		log.Fatal(runErr)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

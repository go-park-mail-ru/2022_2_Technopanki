package main

import (
	"HeadHunter/internal/network"
	"github.com/spf13/viper"
	"log"
)

// @title Jobflow API
// @version 1.0
// @description Swagger API for Golang Project Jobflow.

// @host      95.163.208.72:8080
// @BasePath  /

func main() {
	if configErr := initConfig(); configErr != nil {
		log.Fatal(configErr.Error())
	}

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

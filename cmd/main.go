package main

import (
	"log"

	"github.com/Fastocher/restapp"
	"github.com/Fastocher/restapp/pkg/handler"
	"github.com/Fastocher/restapp/pkg/repository"
	"github.com/Fastocher/restapp/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Initialization falied : %s", err.Error())
	}

	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	srv := new(restapp.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error on start server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

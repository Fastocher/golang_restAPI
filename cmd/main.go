package main

import (
	"log"

	"github.com/Fastocher/restapp"
	"github.com/Fastocher/restapp/pkg/handler"
	"github.com/Fastocher/restapp/pkg/repository"
	"github.com/Fastocher/restapp/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	srv := new(restapp.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error on start server: %s", err.Error())
	}
}

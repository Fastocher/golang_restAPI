package main

import (
	"log"

	"github.com/Fastocher/restapp"
)

func main() {
	srv := new(restapp.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("Error on start server: %s", err.Error())
	}
}

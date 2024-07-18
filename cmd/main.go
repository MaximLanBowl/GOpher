package main

import (
	"log"

	todo "github.com/MaximLanBowl/CRUDGO"
	"github.com/MaximLanBowl/CRUDGO/pkg/handler"
)


func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running server - %s", err.Error())
	}
}
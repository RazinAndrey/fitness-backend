package main

import (
	"log"
	fitnessbackend "github.com/RazinAndrey/fitness-backend"
	"github.com/RazinAndrey/fitness-backend/pkg/handler"
	"github.com/RazinAndrey/fitness-backend/pkg/repository"
	"github.com/RazinAndrey/fitness-backend/pkg/service"
)

func main() {
	// наши зависимости 
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	// иницилизация экземпляра сервера
	srv := new(fitnessbackend.Server)
	// запуск сервера
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

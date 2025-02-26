package main

import (
	fitnessbackend "github.com/RazinAndrey/fitness-backend"
	"github.com/RazinAndrey/fitness-backend/internal/api/v1"
	"github.com/RazinAndrey/fitness-backend/internal/repository"
	"github.com/RazinAndrey/fitness-backend/internal/service"
	"log"
)

func main() {
	// наши зависимости
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := v1.NewHandler(service)

	// иницилизация экземпляра сервера
	srv := new(fitnessbackend.Server)
	// запуск сервера
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

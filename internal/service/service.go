package service

import (
	"github.com/RazinAndrey/fitness-backend/internal/repository"
)

type Authorization interface {
}

type Trainer interface {
}

// собираем сервисы в одном месте
type Service struct {
	Authorization
	Trainer
}

// конструктор
func NewService(repos *repository.Repository) *Service {
	return &Service{}
}

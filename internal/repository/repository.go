package repository

type Authorization interface {
}

type Trainer interface {
}

type Repository struct {
	Authorization
	Trainer
}

// конструктор
func NewRepository() *Repository {
	return &Repository{}
}

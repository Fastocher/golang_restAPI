package repository

type Authorization interface {
}

type Users interface {
}

type Repository struct {
	Authorization
	Users
}

func NewRepository() *Repository {
	return &Repository{}
}

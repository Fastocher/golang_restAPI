package repository

import (
	"github.com/Fastocher/restapp"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user restapp.User) (int, error)
	GetUser(username, password string) (restapp.User, error)
}

type Users interface {
}

type Repository struct {
	Authorization
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}

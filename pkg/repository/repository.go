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

type Message interface {
	CreateMessage(userId int, message restapp.Message) (int, error)
	GetAll(userId int) ([]restapp.Message, error)
	GetById(userId, messageId int) (restapp.Message, error)
}

type Repository struct {
	Authorization
	Users
	Message
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Message:       NewMessagesPostgres(db),
	}
}

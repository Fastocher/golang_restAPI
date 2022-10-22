package service

import (
	"github.com/Fastocher/restapp"
	"github.com/Fastocher/restapp/pkg/repository"
)

type Authorization interface {
	CreateUser(user restapp.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Users interface {
}

type Message interface {
	CreateMessage(userId int, message restapp.Message) (int, error)
	GetAll(userId int) ([]restapp.Message, error)
	GetById(userId, messageId int) (restapp.Message, error)
}

type Service struct {
	Authorization
	Users
	Message
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Message:       NewMessageService(repos.Message),
	}
}

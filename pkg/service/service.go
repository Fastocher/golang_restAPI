package service

import (
	"github.com/Fastocher/restapp"
	"github.com/Fastocher/restapp/pkg/repository"
)

type Authorization interface {
	CreateUser(user restapp.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Users interface {
}

type Service struct {
	Authorization
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

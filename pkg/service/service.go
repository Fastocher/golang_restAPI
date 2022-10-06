package service

import "github.com/Fastocher/restapp/pkg/repository"

type Authorization interface {
}

type Users interface {
}

type Service struct {
	Authorization
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}

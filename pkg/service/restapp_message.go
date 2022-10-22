package service

import (
	"github.com/Fastocher/restapp"
	"github.com/Fastocher/restapp/pkg/repository"
)

type MessageService struct {
	repo repository.Message
}

// конструктор
func NewMessageService(repo repository.Message) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) CreateMessage(userId int, message restapp.Message) (int, error) {
	return s.repo.CreateMessage(userId, message)
}

func (s *MessageService) GetAll(userId int) ([]restapp.Message, error) {
	return s.repo.GetAll(userId)
}

func (s *MessageService) GetById(userId, messageId int) (restapp.Message, error) {
	return s.repo.GetById(userId, messageId)
}

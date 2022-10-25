package restapp

import "errors"

type UsersMessages struct {
	Id        int
	UserId    int
	MessageId int
}

// теги db для возможности выборки из базы
type Message struct {
	Id      int    `json:"id" db:"id"`
	Message string `json:"message" db:"message" binding:"required"`
}

type UpdateMessageInput struct {
	Message *string `json:"message"`
}

func (i UpdateMessageInput) Validate() error {
	if i.Message == nil {
		return errors.New("no values for update")
	}
	return nil
}

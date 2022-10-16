package repository

import (
	"fmt"

	"github.com/Fastocher/restapp"
	"github.com/jmoiron/sqlx"
)

type MessagesPostgres struct {
	db *sqlx.DB
}

func NewMessagesPostgres(db *sqlx.DB) *MessagesPostgres {
	return &MessagesPostgres{db: db}
}

func (r *MessagesPostgres) CreateMessage(userId int, message restapp.Message) (int, error) {
	//начало транзакции
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	// вставка в таблицу MESSAGE
	createMessageQuery := fmt.Sprintf("INSERT INTO %s (message) VALUES ($1) RETURNING id", messagesTable)
	row := tx.QueryRow(createMessageQuery, message.Message)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	// вставка в таблицу Users_Messages
	createUsersMessagesQuery := fmt.Sprintf("INSERT INTO %s(user_id, message_id) VALUES ($1, $2)", usersMessagestable)
	_, err = tx.Exec(createUsersMessagesQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	// ззавершение транзакции. Коммит.
	return id, tx.Commit()
}

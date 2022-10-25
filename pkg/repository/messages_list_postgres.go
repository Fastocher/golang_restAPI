package repository

import (
	"fmt"
	"strings"

	"github.com/Fastocher/restapp"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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
	// завершение транзакции. Коммит.
	return id, tx.Commit()
}

// имплементация метода описанного в интерфейсе репозитория
func (r *MessagesPostgres) GetAll(userId int) ([]restapp.Message, error) {
	var messages []restapp.Message
	// выбираю значения по inner join, тоесть те значения которые имеют
	// одинаковые значения в обоих таблицах
	// все записи из таблицы message которые есть в табилце users messages
	// и при этом связаны с id пользователя
	query := fmt.Sprintf("SELECT m.id,m.message FROM %s m INNER JOIN %s ul on m.id = ul.message_id WHERE ul.user_id = $1",
		messagesTable, usersMessagestable)
	err := r.db.Select(&messages, query, userId)

	return messages, err

}

func (r *MessagesPostgres) GetById(userId, messageId int) (restapp.Message, error) {
	var message restapp.Message

	query := fmt.Sprintf(`SELECT m.id,m.message FROM %s m INNER JOIN %s um 
	on m.id = um.message_id WHERE um.user_id = $1 AND um.message_id = $2`, messagesTable, usersMessagestable)
	err := r.db.Get(&message, query, userId, messageId)

	return message, err
}

func (r *MessagesPostgres) Delete(userId, messageId int) error {
	query := fmt.Sprintf("DELETE FROM %s m USING %s um WHERE m.id = um.message_id AND um.user_id=$1 AND um.message_id=$2",
		messagesTable, usersMessagestable)

	_, err := r.db.Exec(query, userId, messageId)

	return err
}

func (r *MessagesPostgres) Update(userId, messageId int, input restapp.UpdateMessageInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Message != nil {
		setValues = append(setValues, fmt.Sprintf("message=$%d", argId))
		args = append(args, *input.Message)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s m SET %s FROM %s um WHERE m.id = um.message_id AND um.message_id=$%d AND um.user_id=$%d",
		messagesTable, setQuery, usersMessagestable, argId, argId+1)

	args = append(args, messageId, userId)

	logrus.Debugf("update query: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)

	return err

}

package repository

import (
	"fmt"

	"github.com/Fastocher/restapp"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user restapp.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2 ,$3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// метод репозитория для получения пользотеля по логину и паролю
func (r *AuthPostgres) GetUser(username, password string) (restapp.User, error) {
	var user restapp.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	//метод базы данных GET и передаём указатель на структуру &user
	err := r.db.Get(&user, query, username, password)

	return user, err
}

package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Users interface {
}

type Repository struct {
	Authorization
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}

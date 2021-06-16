package repository

import (
	"github.com/jmoiron/sqlx"
	"restApi"
)

type Authorization interface {
	CreateUser(user restApi.User) (int, error)
	GetUser(username, password string) (restApi.User, error)
}

type TodoList interface {
	CreateList(userid int, list restApi.TodoList) (int, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}

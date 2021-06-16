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
	GetAllLists(userid int) ([]restApi.TodoList, error)
	GetListById(userid int, listid int) (restApi.TodoList, error)
	DeleteListById(id int, listid int) error
	UpdateList(userId int, listid int, update restApi.UpdateListInput) error
}

type TodoItem interface {
	CreateItem(listId int, input restApi.TodoItem) (int, error)
	GetAllItems(userid, listId int) ([]restApi.TodoItem, error)
	GetItemById(userId, id int) (restApi.TodoItem, error)
	DeleteItem(userId, id int) error
	UpdateItem(userId, id int, update restApi.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}

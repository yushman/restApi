package service

import (
	"restApi"
	"restApi/pkg/repository"
)

type Authorization interface {
	CreateUser(user restApi.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	CreateList(userid int, list restApi.TodoList) (int, error)
	GetAllLists(userid int) ([]restApi.TodoList, error)
	GetListById(userid int, listid int) (restApi.TodoList, error)
	DeleteListById(userId int, listid int) error
	UpdateList(userId int, listid int, update restApi.UpdateListInput) error
}

type TodoItem interface {
	CreateItem(userId int, listId int, input restApi.TodoItem) (int, error)
	GetAllItems(userId int, listId int) ([]restApi.TodoItem, error)
	GetItemById(userId int, id int) (restApi.TodoItem, error)
	DeleteItem(userId, id int) error
	UpdateItem(userId, id int, update restApi.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
		TodoItem:      NewTodoItemService(repo.TodoItem, repo.TodoList),
	}
}

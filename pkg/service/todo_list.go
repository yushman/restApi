package service

import (
	"restApi"
	"restApi/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (t *TodoListService) CreateList(userid int, list restApi.TodoList) (int, error) {
	return t.repo.CreateList(userid, list)
}

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

func (t *TodoListService) GetAllLists(userid int) ([]restApi.TodoList, error) {
	return t.repo.GetAllLists(userid)
}

func (t *TodoListService) GetListById(userid int, listid int) (restApi.TodoList, error) {
	return t.repo.GetListById(userid, listid)
}

func (t *TodoListService) DeleteListById(userId int, listid int) error {
	return t.repo.DeleteListById(userId, listid)
}

func (t *TodoListService) UpdateList(userId int, listid int, update restApi.UpdateListInput) error {
	if err := update.Validate(); err != nil {
		return err
	}
	return t.repo.UpdateList(userId, listid, update)
}

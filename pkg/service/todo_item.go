package service

import (
	"restApi"
	"restApi/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listrepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listrepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listrepo: listrepo}
}

func (t *TodoItemService) CreateItem(userId int, listId int, input restApi.TodoItem) (int, error) {
	if _, err := t.listrepo.GetListById(userId, listId); err != nil {
		return 0, err
	}
	return t.repo.CreateItem(listId, input)
}

func (t *TodoItemService) GetAllItems(userId int, listId int) ([]restApi.TodoItem, error) {
	return t.repo.GetAllItems(userId, listId)
}

func (t *TodoItemService) GetItemById(userId int, id int) (restApi.TodoItem, error) {
	return t.repo.GetItemById(userId, id)
}

func (t *TodoItemService) DeleteItem(userId, id int) error {
	return t.repo.DeleteItem(userId, id)
}

func (t *TodoItemService) UpdateItem(userId, id int, update restApi.UpdateItemInput) error {
	if err := update.Validate(); err != nil {
		return err
	}
	return t.repo.UpdateItem(userId, id, update)
}

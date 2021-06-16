package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"restApi"
	"strings"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (t *TodoItemPostgres) CreateItem(listId int, input restApi.TodoItem) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createItemQuery, input.Title, input.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createListsItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2) RETURNING id", listsItemsTable)
	_, err = tx.Exec(createListsItemsQuery, listId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (t *TodoItemPostgres) GetAllItems(userid, listid int) ([]restApi.TodoItem, error) {
	var items []restApi.TodoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti "+
		"INNER JOIN %s li on ti.id = li.item_id "+
		"INNER JOIN %s ul on ul.list_id = li.list_id "+
		"WHERE li.list_id = $1 "+
		"AND ul.user_id = $2", todoItemsTable, listsItemsTable, usersListsTable)
	err := t.db.Select(&items, query, listid, userid)
	return items, err
}

func (t *TodoItemPostgres) GetItemById(userId, id int) (restApi.TodoItem, error) {
	var item restApi.TodoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti "+
		"INNER JOIN %s li on ti.id = li.item_id "+
		"INNER JOIN %s ul on li.list_id = ul.list_id "+
		"WHERE ti.id = $1 "+
		"AND ul.user_id = $2", todoItemsTable, listsItemsTable, usersListsTable)
	err := t.db.Get(&item, query, id, userId)
	return item, err
}

func (t *TodoItemPostgres) DeleteItem(userId, id int) error {
	query := fmt.Sprintf("DELETE from %s ti USING %s li, %s ul "+
		"WHERE ti.id = li.item_id "+
		"AND li.list_id = ul.list_id "+
		"AND ul.user_id = $1 "+
		"AND ti.id = $2", todoItemsTable, listsItemsTable, usersListsTable)
	_, err := t.db.Exec(query, userId, id)
	return err
}

func (t *TodoItemPostgres) UpdateItem(userId, id int, update restApi.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argsId := 1

	if update.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argsId))
		args = append(args, *update.Title)
		argsId++
	}
	if update.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argsId))
		args = append(args, *update.Description)
		argsId++
	}
	if update.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argsId))
		args = append(args, *update.Done)
		argsId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s ti SET %s FROM %s li, %s ul "+
		"WHERE ti.id = li.item_id "+
		"AND li.list_id = ul.list_id "+
		"AND ul.user_id = $%d "+
		"AND ti.id = $%d", todoItemsTable, setQuery, listsItemsTable, usersListsTable, argsId, argsId+1)

	args = append(args, userId, id)
	fmt.Printf("%s\n", query)
	fmt.Printf("%s\n", args)
	_, err := t.db.Exec(query, args...)
	return err
}

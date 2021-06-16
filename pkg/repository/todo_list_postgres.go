package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"restApi"
	"strings"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (t *TodoListPostgres) CreateList(userid int, list restApi.TodoList) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListsQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2) RETURNING id", usersListsTable)
	_, err = tx.Exec(createUsersListsQuery, userid, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (t *TodoListPostgres) GetAllLists(userid int) ([]restApi.TodoList, error) {
	var lists []restApi.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl "+
		"INNER JOIN %s ul on tl.id = ul.list_id "+
		"WHERE ul.user_id=$1", todoListsTable, usersListsTable)
	err := t.db.Select(&lists, query, userid)
	return lists, err
}

func (t *TodoListPostgres) GetListById(userid int, listid int) (restApi.TodoList, error) {
	var list restApi.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl "+
		"INNER JOIN %s ul on tl.id = ul.list_id "+
		"WHERE ul.user_id = $1 "+
		"AND tl.id = $2", todoListsTable, usersListsTable)
	err := t.db.Get(&list, query, userid, listid)
	return list, err
}

func (t *TodoListPostgres) DeleteListById(userId int, listid int) error {
	query := fmt.Sprintf("DELETE from %s tl USING %s ul "+
		"WHERE tl.id = ul.list_id "+
		"AND ul.user_id = $1 "+
		"AND ul.list_id = $2", todoListsTable, usersListsTable)
	_, err := t.db.Exec(query, userId, listid)
	return err
}

func (t *TodoListPostgres) UpdateList(userId int, listid int, update restApi.UpdateListInput) error {
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
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul "+
		"WHERE tl.id = ul.list_id "+
		"AND ul.list_id = $%d "+
		"AND ul.user_id = $%d", todoListsTable, setQuery, usersListsTable, argsId, argsId+1)

	args = append(args, listid, userId)
	_, err := t.db.Exec(query, args...)
	return err
}

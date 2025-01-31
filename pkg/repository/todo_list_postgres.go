package repository

import (
	"fmt"

	"github.com/MyNameIsWhaaat/todo-app"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error){
	logrus.Infof("Create() called with userId: %d, list title: %s, description: %s", userId, list.Title, list.Description)

    tx, err := r.db.Begin()
    if err != nil {
        logrus.Errorf("failed to begin transaction: %s", err.Error())
        return 0, err
    }

    var id int
    createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
    row := tx.QueryRow(createListQuery, list.Title, list.Description)
    if err := row.Scan(&id); err != nil {
        tx.Rollback()
        logrus.Errorf("failed to scan id: %s", err.Error())
        return 0, err
    }

    createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
    _, err = tx.Exec(createUsersListQuery, userId, id)
    if err != nil {
        tx.Rollback()
        logrus.Errorf("failed to execute users list query: %s", err.Error())
        return 0, err
    }

    if err := tx.Commit(); err != nil {
		logrus.Errorf("failed to commit transaction: %s", err.Error())
		return 0, err
	}
	logrus.Info("Transaction committed successfully, ID:", id)

    return id, nil
}

func (r *TodoListPostgres) 	GetAll(userId int) ([]todo.TodoList, error){
    var lists []todo.TodoList
    query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
        todoListsTable, usersListsTable)
    err := r.db.Select(&lists, query, userId)

    return lists, err
}

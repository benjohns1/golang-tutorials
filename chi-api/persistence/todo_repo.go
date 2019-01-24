package persistence

import (
	"database/sql"
	"fmt"

	"github.com/benjohns1/golang-tutorials/chi-api/todo"
)

type TodoRepository struct {
	conn *sql.DB
}

type TodoUnitOfWork struct {
	todoRepo *TodoRepository
	conn     *sql.DB
}

func NewTodoUnitOfWork(conn *sql.DB) TodoUnitOfWork {
	return TodoUnitOfWork{
		todoRepo: &TodoRepository{conn},
		conn:     conn,
	}
}

func (t *TodoRepository) Add(todo *todo.Todo) error {
	return fmt.Errorf("NOT IMPLEMENTED Unable to add todo")
}

func (t *TodoRepository) Remove(todo *todo.Todo) error {
	return fmt.Errorf("NOT IMPLEMENTED Unable to remove todo")
}

// Get retrieve a single Todo
func (t *TodoRepository) Get(todoID int) (*todo.Todo, error) {

	query, err := t.conn.Prepare("SELECT slug, title, body FROM todo WHERE id=$1")
	if err != nil {
		return nil, fmt.Errorf("Could not query database: %v", err)
	}
	rows, err := query.Query(todoID)
	if err != nil {
		return nil, fmt.Errorf("Error querying database: %v", err)
	}

	var todo todo.Todo
	if rows.Next() {
		err = rows.Scan(&todo.Slug, &todo.Title, &todo.Body)
		if err != nil {
			return nil, fmt.Errorf("Error parsing database data: %v", err)
		}
	} else {
		return nil, fmt.Errorf("No todo with ID %v", todoID)
	}

	return &todo, nil
}

func (t *TodoRepository) GetAll() ([]*todo.Todo, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED Unable to get all todos")
}

func (u TodoUnitOfWork) Complete() error {
	return fmt.Errorf("NOT IMPLEMENTED Unable to complete")
}

func (u TodoUnitOfWork) Todo() todo.Repository {
	return u.todoRepo
}

package postgresdb

import (
	"database/sql"
	"fmt"

	"github.com/benjohns1/golang-tutorials/chi-api/core"
)

// TodoRepository provides access to Todos
type TodoRepository struct {
	conn *sql.DB
}

// Add creates a new Todo
func (t *TodoRepository) Add(todo *core.Todo) (todoID int, err error) {

	insert, err := t.conn.Prepare("INSERT INTO todo (slug, title, body) VALUES ($1, $2, $3) RETURNING id;")
	if err != nil {
		return 0, fmt.Errorf("Error compiling database insert statement: %v", err)
	}
	err = insert.QueryRow(todo.Slug, todo.Title, todo.Body).Scan(&todoID)
	if err != nil {
		return 0, fmt.Errorf("Error inserting row into database: %v", err)
	}

	return todoID, nil
}

// Remove remotes a single Todo by its ID
func (t *TodoRepository) Remove(todoID int) error {
	delete, err := t.conn.Prepare("DELETE FROM todo WHERE id=$1")
	if err != nil {
		return fmt.Errorf("Error compiling database delete statement: %v", err)
	}

	res, err := delete.Exec(todoID)
	if err != nil {
		return fmt.Errorf("Error deleting from database: %v", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error deleting todo %v: ", err)
	}
	if count == 1 {
		return nil
	}
	if count <= 0 {
		return fmt.Errorf("No todos deleted, todoID %v not found", todoID)
	}
	return fmt.Errorf("Multiple todos deleted, count: %v", count)
}

// Get retrieves a single Todo
func (t *TodoRepository) Get(todoID int) (*core.Todo, error) {

	query, err := t.conn.Prepare("SELECT slug, title, body FROM todo WHERE id=$1")
	if err != nil {
		return nil, fmt.Errorf("Error compiling database query: %v", err)
	}
	rows, err := query.Query(todoID)
	if err != nil {
		return nil, fmt.Errorf("Error querying database: %v", err)
	}

	var todo core.Todo
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

// GetAll gets all Todos
func (t *TodoRepository) GetAll() ([]*core.Todo, error) {

	query, err := t.conn.Prepare("SELECT slug, title, body FROM todo")
	if err != nil {
		return nil, fmt.Errorf("Error compiling database query: %v", err)
	}
	rows, err := query.Query()
	if err != nil {
		return nil, fmt.Errorf("Error querying database: %v", err)
	}

	var todos []*core.Todo
	for rows.Next() {
		var todo core.Todo
		err = rows.Scan(&todo.Slug, &todo.Title, &todo.Body)
		if err != nil {
			return nil, fmt.Errorf("Error parsing database data: %v", err)
		}
		todos = append(todos, &todo)
	}

	return todos, nil
}

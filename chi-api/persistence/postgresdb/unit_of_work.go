package postgresdb

import (
	"database/sql"
	"fmt"

	"github.com/benjohns1/golang-tutorials/chi-api/core"
)

// UnitOfWork concrete implementation of a unit of work
type UnitOfWork struct {
	todoRepo *TodoRepository
	conn     *sql.DB
}

// NewUnitOfWork Constructs a new unit of work object from a database connection
func NewUnitOfWork(conn *sql.DB) UnitOfWork {
	return UnitOfWork{
		todoRepo: &TodoRepository{conn},
		conn:     conn,
	}
}

// Complete applies all in-memory changes to the database
func (u UnitOfWork) Complete() error {
	panic(fmt.Errorf("NOT IMPLEMENTED"))
}

// Todo provides access to TodoRepository
func (u UnitOfWork) Todo() core.TodoRepository {
	return u.todoRepo
}

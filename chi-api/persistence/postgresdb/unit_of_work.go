package postgresdb

import (
	"database/sql"
	"fmt"

	"github.com/benjohns1/golang-tutorials/chi-api/core"
)

type UnitOfWork struct {
	todoRepo *TodoRepository
	conn     *sql.DB
}

func NewUnitOfWork(conn *sql.DB) UnitOfWork {
	return UnitOfWork{
		todoRepo: &TodoRepository{conn},
		conn:     conn,
	}
}

func (u UnitOfWork) Complete() error {
	return fmt.Errorf("NOT IMPLEMENTED Unable to complete")
}

func (u UnitOfWork) Todo() core.TodoRepository {
	return u.todoRepo
}

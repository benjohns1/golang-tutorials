package core

// Todo domain entity contains todo data
type Todo struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// TodoRepository wraps persistent actions on todos
type TodoRepository interface {
	Add(*Todo) (todoID int, err error)
	Remove(todoID int) error
	Get(todoID int) (*Todo, error)
	GetAll() ([]*Todo, error)
}

// GetTodo gets a single todo
func GetTodo(u UnitOfWork, todoID int) (*Todo, error) {
	return u.Todo().Get(todoID)
}

func DeleteTodo(u UnitOfWork, todoID int) error {
	return u.Todo().Remove(todoID)
}

func AddTodo(u UnitOfWork, todo *Todo) (todoID int, err error) {
	return u.Todo().Add(todo)
}

func GetAllTodos(u UnitOfWork) ([]*Todo, error) {
	return u.Todo().GetAll()
}

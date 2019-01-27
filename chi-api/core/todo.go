package core

// Todo domain entity contains todo data
type Todo struct {
	ID    int
	Slug  string
	Title string
	Body  string
}

// TodoRepository wraps persistent actions on todos
type TodoRepository interface {
	Add(*Todo) (todoID int, err error)
	Remove(todoID int) error
	Get(todoID int) (*Todo, error)
	GetAll() ([]*Todo, error)
}

// GetTodo get a single todo, given its ID
func GetTodo(u UnitOfWork, todoID int) (*Todo, error) {
	return u.Todo().Get(todoID)
}

// DeleteTodo delete a single todo, given its ID
func DeleteTodo(u UnitOfWork, todoID int) error {
	return u.Todo().Remove(todoID)
}

// AddTodo add a single todo
func AddTodo(u UnitOfWork, todo *Todo) (todoID int, err error) {
	return u.Todo().Add(todo)
}

// GetAllTodos retrieve all stored todos
func GetAllTodos(u UnitOfWork) ([]*Todo, error) {
	return u.Todo().GetAll()
}

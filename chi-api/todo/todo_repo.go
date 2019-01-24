package todo

// Repository wraps persistent actions on todos
type Repository interface {
	Add(*Todo) error
	Remove(*Todo) error
	Get(todoID int) (*Todo, error)
	GetAll() ([]*Todo, error)
}

// UnitOfWork manages persistent repositories and updates for todos
type UnitOfWork interface {
	Todo() Repository
	Complete() error
}

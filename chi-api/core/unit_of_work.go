package core

// UnitOfWork manages persistent repositories and updates for todos
type UnitOfWork interface {
	Todo() TodoRepository
	Complete() error
}

package todo

import "errors"

var (
	ErrTodoNotFound    = errors.New("todo not found")
	ErrInvalidTodoData = errors.New("invalid todo data")
)

package todo

// CreateTodoRequest - запрос на создание Todo
type CreateTodoRequest struct {
	Title       string `json:"title" validate:"required, min=2"`
	Description string `json: "description"`
}

// UpdateTodoRequest - запрос на обновление Todo
type UpdateTodoRequest struct {
	Title       string `json:"title" validate:"omitempty, min=2"`
	Description string `json:="description"`
	Completed   bool   `json:"completed"`
}

// TodoResponse - ответ с данными Todo
type TodoResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json: "created_at"`
	UpdatedAt   string `json:"updated_at"`
}

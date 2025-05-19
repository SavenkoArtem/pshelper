package todo

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog"
)

type TodoHandlerDeps struct {
	CustomLogger *zerolog.Logger
}

type TodoHandler struct {
	customLogger *zerolog.Logger
}

func NewTodoHandler(router *http.ServeMux, deps TodoHandlerDeps) {
	handler := &TodoHandler{
		customLogger: deps.CustomLogger,
	}
	router.HandleFunc("GET /todo", handler.GetAll())
}

func (h *TodoHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("sdf")
	}
}

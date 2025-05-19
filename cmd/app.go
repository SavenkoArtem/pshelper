package main

import (
	"net/http"

	"github.com/SavenkoArtem/pshelper/configs"
	todo "github.com/SavenkoArtem/pshelper/internal/todo/handler"
	"github.com/SavenkoArtem/pshelper/pkg/database"
	"github.com/SavenkoArtem/pshelper/pkg/logger"
)

func App() (http.Handler, string) {
	cfg := configs.LoadConfig()
	router := http.NewServeMux()
	customLogger := logger.InitLogger(&cfg.Logger)
	dbpool := database.CreateDbPool(cfg, customLogger)
	defer dbpool.Close()

	// Repositories

	// Services

	// Handlers
	todo.NewTodoHandler(router, todo.TodoHandlerDeps{
		CustomLogger: customLogger,
	})

	return router, cfg.Server.Port
}

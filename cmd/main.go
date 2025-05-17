package main

import (
	"fmt"
	"net/http"

	"github.com/SavenkoArtem/pshelper/internal/home" // Импортируем обработчик домашней страницы
)

func main() {
	// Регистрируем обработчики
	http.HandleFunc("/", home.NewHandler)

	// Настраиваем порт
	port := "3000"
	fmt.Printf("Server started %s\n", port)
	fmt.Println("Go to the http://localhost:" + port)

	// Запускаем сервер
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error start server: %s\n", err)
	}
}

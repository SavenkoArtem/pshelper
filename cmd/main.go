package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler, port := App()

	server := http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	fmt.Println("Server is listening on port " + port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

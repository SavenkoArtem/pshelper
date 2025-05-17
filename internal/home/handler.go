package home

import (
	"fmt"
	"net/http"
)

func NewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi! This is the first Go server!")
}

package handlers

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

// this is a test handler to know if everything work

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "everything is good\n")
}

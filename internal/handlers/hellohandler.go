package handlers

import (
	"fmt"
	"net/http"
	"sync"
)

type helloHandler struct {
	mu sync.Mutex
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "everything is good\n")
}

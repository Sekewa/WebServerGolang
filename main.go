package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	//"time"
)

type countHandler struct {
	mu sync.Mutex
	n  int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

func Exit(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "BYE !!")
	log.Fatal("BYE !!")
}

func main() {
	http.Handle("/count", new(countHandler))

	http.HandleFunc("/exit", Exit)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

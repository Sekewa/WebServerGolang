package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func Exit(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "BYE !!")
	log.Fatal("BYE !!")
}

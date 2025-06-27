package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Health struct {
	HealthCheck string
}

// this handler will return if the integrities of files are all good

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	health := Health{"OK"}

	h, err := json.Marshal(health)

	if err != nil {
		log.Println("Error : couldn't parse health response")
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(h)
}

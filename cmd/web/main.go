package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Sekewa/WebServerGolang/internal/utils"
)

/*
Exemple supplementaire
*/
func Exit(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "BYE !!")
	log.Fatal("BYE !!")
}

func main() {
	params := utils.ArgsParser(os.Args)

	logger := utils.NewLogger()

	mux := http.NewServeMux()

	mux.Handle("/hello", new(handlers.helloHandler))
	mux.HandleFunc("/exit", Exit)

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/", fs)

	loggedMux := utils.LoggingMiddleware(mux, logger)

	srv := &http.Server{
		Addr:         params["port"],
		Handler:      loggedMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("Start of serveur on http://localhost", params["port"])

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server stopped : %v", err)
	}
}

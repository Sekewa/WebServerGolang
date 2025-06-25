package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Sekewa/WebServerGolang/internal/handlers"
	"github.com/Sekewa/WebServerGolang/internal/utils"
)

func main() {
	params := utils.ArgsParser(os.Args)

	logger := utils.NewLogger()

	mux := http.NewServeMux()

	// handler avec une struct
	mux.Handle("/hello", new(handlers.HelloHandler))
	// handler avec une fonction
	mux.HandleFunc("/exit", handlers.Exit)

	// handler avec un dossier
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

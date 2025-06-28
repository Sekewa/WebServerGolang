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

	// test handler that will return hello
	mux.Handle("/hello", new(handlers.HelloHandler))
	// handler that will return stat of this web server
	mux.HandleFunc("/health", handlers.HealthCheck)

	// handler of the example web site or static one
	fs := http.FileServer(utils.SearchStatic())
	mux.Handle("/", fs)

	// logger that will log every interaction with the web server
	loggedMux := utils.LoggingMiddleware(mux, logger)

	// server struct
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

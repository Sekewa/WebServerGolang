package utils

import (
	"log"
	"net/http"
	"time"
)

func NewLogger() *log.Logger {
	return log.Default()
}

// its a middleware who log every HTTP request
func LoggingMiddleware(next http.Handler, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		elapsed := time.Since(start)

		logger.Printf("%s %s -> %v", r.Method, r.URL.Path, elapsed)
	})
}

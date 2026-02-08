// Package middleware provides cross-cutting HTTP request handling utilities.
// Middleware functions wrap the HTTP request flow to add functionality like logging.
package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// Logger is an HTTP middleware that logs information about each request.
// It records the HTTP method, request path, and execution time for debugging and monitoring.
// Example output: "GET /pages/123 2.345ms"
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Println(r.Method, r.URL.Path, time.Since(start))
	})
}

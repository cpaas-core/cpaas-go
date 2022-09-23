package main

import (
	"fmt"
	"net/http"
	"os"
)

func handleHostname(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%s\n", hostname)
}

func Methods(next http.HandlerFunc, methods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		validMethod := false
		for _, method := range methods {
			if r.Method == method {
				validMethod = true
			}
		}

		if !validMethod {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		next(w, r)
	}
}

func HandleFunc(method string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method && method != "*" {
			http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
			return
		}

		next(w, r)
	}
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging")
		next(w, r)
	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logged in!")
		next(w, r)
	}
}

func main() {
	// http.HandleFunc("/hostname", Methods(handleHostname, "GET", "POST"))
	http.HandleFunc("/hostname", HandleFunc("GET", handleHostname))

	fmt.Println("Running server on http://localhost:5000")
	http.ListenAndServe(":5000", AuthMiddleware(LoggingMiddleware(http.DefaultServeMux.ServeHTTP)))
}

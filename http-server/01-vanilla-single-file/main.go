package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Write here your routes and handlers
	fmt.Println("Running server on http://localhost:5000")
	http.ListenAndServe(":5000", nil)
}

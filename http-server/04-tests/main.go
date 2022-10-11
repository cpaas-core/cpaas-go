package main

import (
	"application/app"
	"fmt"
	"net/http"
)

func main() {
	server := app.NewServer()

	fmt.Println("Starting application on http://localhost:5000")
	http.ListenAndServe(":5000", server)
}

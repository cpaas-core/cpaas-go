package main

import (
	"net/http"
)

func main() {
	handler := app.NewHandler()
	http.ListenAndServe(":5000", handler)
}

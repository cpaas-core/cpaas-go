package main

import (
	"net/http"

	"workspace/app"
)

func main() {
	handler := app.NewHandler()
	http.ListenAndServe(":5000", handler)
}

package main

import (
	"net/http"
        "httpmod/app"
)

func main() {

	handler := app.NewHandler()
	http.ListenAndServe(":5000", handler)

}

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

func main() {
	http.HandleFunc("/hostname", handleHostname)

	fmt.Println("Running server on http://localhost:5000")
	http.ListenAndServe(":5000", nil)
}

package app

import (
	"fmt"
	"net/http"
	"os"
)

func HandleHostname(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%s", hostname)
}

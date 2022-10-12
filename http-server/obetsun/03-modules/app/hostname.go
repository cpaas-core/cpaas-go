package app

import (
	"fmt"
	"net/http"
	"os"
)

func HostnameHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Fprintf(w, "Hostname: %s", hostname)

}

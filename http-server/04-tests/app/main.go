package app

import "net/http"

func NewServer() *http.ServeMux {
	serveMux := http.ServeMux{}

	serveMux.HandleFunc("/hostname", HandleHostname)

	return &serveMux
}

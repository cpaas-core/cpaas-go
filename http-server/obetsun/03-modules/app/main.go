package app

import (
	"net/http"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/hash", HashHandler)
	mux.HandleFunc("/hostname", HostnameHandler)
	mux.HandleFunc("/headers", HeadersHandler)
	mux.HandleFunc("/", NslookupHandler)

	return mux
}

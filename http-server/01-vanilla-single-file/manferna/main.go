package main

// Sources:
// https://medium.com/rungo/creating-a-simple-hello-world-http-server-in-go-31c7fd70466e
// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
// https://gobyexample.com/http-servers

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
)

func test(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "¡Hola! this is a test")
}

func headers(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		fmt.Fprintf(w, "No no no, just GET method is supported!")
		return
	}

	for name, headers := range req.Header {
		for _, header := range headers {
			fmt.Fprintf(w, "%s: %s", name, header)
		}
	}
}

func hostname(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		fmt.Fprintf(w, "No no no, just GET method is supported!")
		return
	}
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
	fmt.Fprintf(w, "Hostname: %s", hostname)

}

func nslookup(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		fmt.Fprintf(w, "No no no, just GET method is supported!")
		return
	}

	_, hostname := path.Split(req.URL.Path)

	ipaddress, err := net.LookupIP(hostname)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
	fmt.Fprintf(w, "IP address for hostname %s -> %s", hostname, ipaddress[0])
}

func hash(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		fmt.Fprintf(w, "No no no, just GET method is supported!")
		return
	}

	// text and algorithm must exist
	query := req.URL.Query()
	if query["text"] == nil || query["algorithm"] == nil {
		fmt.Fprint(w, "text or algorithm values are empty!")
		return
	}

	algorithm := query["algorithm"][0]
	text := query["text"][0]
	var hash string
	switch algorithm {
	case "md5":
		hash = fmt.Sprintf("%x", md5.Sum([]byte(string(text))))
	case "SHA256":
		hash = fmt.Sprintf("%x", sha256.Sum256([]byte(text)))
	case "SHA1":
		hash = fmt.Sprintf("%x", sha1.Sum([]byte(text)))
	default:
		hash = ""
	}

	fmt.Fprintf(w, "text: %s, algorithm: %s, hash: %s", text, algorithm, hash)
}

func main() {
	// Write here your routes and handlers
	fmt.Println("Running server on http://localhost:5000")

	http.HandleFunc("/hostname", hostname)
	http.HandleFunc("/nslookup/", nslookup)
	http.HandleFunc("/hash/", hash)
	http.HandleFunc("/header", headers)
	http.HandleFunc("/", test)

	http.ListenAndServe(":5000", nil)
}

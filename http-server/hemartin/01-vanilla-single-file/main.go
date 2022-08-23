package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path"
)

func main() {
	// Write here your routes and handlers
	http.HandleFunc("/hostname", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			fmt.Fprintf(w, "Method '%s' not supported.\n", r.Method)
			return
		}

		hostname, err := os.Hostname()
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "%s\n", hostname)
	})

	http.HandleFunc("/nslookup/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			fmt.Fprintf(w, "Method '%s' not supported.\n", r.Method)
			return
		}

		_, name := path.Split(r.URL.Path)

		ips, err := net.LookupIP(name)
		if err != nil {
			fmt.Fprintf(w, "%s\n", err)
			return
		}

		fmt.Fprintf(w, "%s\n", ips[0])
	})

	http.HandleFunc("/hash/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			fmt.Fprintf(w, "Method '%s' not supported.\n", r.Method)
			return
		}

		query := r.URL.Query()

		if len(query["text"]) == 0 || len(query["algorithm"]) == 0 {
			fmt.Fprint(w, "Both 'text' and 'algorithm' should be provided as an URL query.\n")
			return
		}
		text := query["text"][0]
		algorithm := query["algorithm"][0]
		finalText := ""

		if algorithm == "md5" {
			hash := md5.New()
			io.WriteString(hash, text)
			finalText = fmt.Sprintf("%x", hash.Sum(nil))
		} else if algorithm == "sha256" {
			hash := sha256.New()
			io.WriteString(hash, text)
			finalText = fmt.Sprintf("%x", hash.Sum(nil))
		} else {
			fmt.Fprintf(w, "algorithm '%s' not supported.\n", algorithm)
			return
		}

		fmt.Fprintf(w, "text: %s, algorithm: %s, hash: %s\n", text, algorithm, finalText)
	})

	http.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			fmt.Fprintf(w, "Method '%s' not supported.\n", r.Method)
			return
		}

		fmt.Fprintf(w, "Headers: %s\n", r.Header)
	})

	fmt.Println("Running server on http://localhost:5000")
	http.ListenAndServe(":5000", nil)
}

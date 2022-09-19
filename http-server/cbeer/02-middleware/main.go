package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"strings"
)

func handleFunc(w http.ResponseWriter, req *http.Request, allowedMethod string) error {
	if req.Method == allowedMethod {
		return nil
	}
	errorMessage := fmt.Sprintf("Only %s requests are allowed!", allowedMethod)
  http.Error(w, errorMessage, http.StatusMethodNotAllowed)
	return errors.New(errorMessage)
}

func AllowMethods(w http.ResponseWriter, req *http.Request, allowedMethods []string) error {
	for _, allowed := range allowedMethods {
		if req.Method == allowed {
			return nil
		}
	}
	allowedString := strings.Join(allowedMethods, ", ")
  errorMessage := fmt.Sprintf("Only the following requests are allowed: %s", allowedString)
  http.Error(w, errorMessage, http.StatusMethodNotAllowed)
  return errors.New(errorMessage)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	if err := handleFunc(w, req, http.MethodGet); err != nil {
		return
	}
//  allowedMethods := []string{http.MethodGet, http.MethodPost}
//	if err := AllowMethods(w, req, allowedMethods); err != nil {
//		return
//	}

	fmt.Fprintf(w, "indexHandler\n")
}

func hello(w http.ResponseWriter, req *http.Request) {
	if err := handleFunc(w, req, http.MethodGet); err != nil {
		return
	}

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	if err := handleFunc(w, req, http.MethodGet); err != nil {
		return
	}

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func hostname(w http.ResponseWriter, req *http.Request) {
	if err := handleFunc(w, req, http.MethodGet); err != nil {
		return
	}

	name, err := os.Hostname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "hostname: %s\n", name)
}

func nslookup(w http.ResponseWriter, req *http.Request) {
	if err := handleFunc(w, req, http.MethodGet); err != nil {
		return
	}

	_, hostname := path.Split(req.URL.Path)
	ipaddr, err := net.LookupIP(hostname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "nslookup IP address: %s\n", ipaddr[len(ipaddr)-1])
}

func hash(w http.ResponseWriter, req *http.Request) {
	if err := handleFunc(w, req, http.MethodGet); err != nil {
		return
	}

	req_text := req.URL.Query().Get("text")
	if req_text == "" {
		http.Error(w, "The text query parameter is missing", http.StatusBadRequest)
		return
	}

	req_algorithm := req.URL.Query().Get("algorithm")
	if req_algorithm == "" {
		http.Error(w, "The algorithm query parameter is missing", http.StatusBadRequest)
		return
	}

	hash := "junk"
	if strings.EqualFold("md5", req_algorithm) {
		h := md5.Sum([]byte(req_text))
		hash = fmt.Sprintf("%x", h)
	} else if strings.EqualFold("sha1", req_algorithm) {
		h := sha1.Sum([]byte(req_text))
		hash = fmt.Sprintf("%x", h)
	} else if strings.EqualFold("sha256", req_algorithm) {
		h := sha256.Sum256([]byte(req_text))
		hash = fmt.Sprintf("%x", h)
	} else {
		errmsg := fmt.Sprintf("The %s algorithm is not supported", req_algorithm)
		http.Error(w, errmsg, http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "{algorithm: %s,text: %s,hash: %s}\n",
		          req_algorithm, req_text, hash)
}


func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/hostname", hostname)
	http.HandleFunc("/nslookup/", nslookup)
	http.HandleFunc("/hash/", hash)

	fmt.Println("Running server on http://localhost:5000")
	http.ListenAndServe(":5000", nil)
}

package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	addr := ":5000"

	mux := http.NewServeMux()

	mux.HandleFunc("/hash", HashHandler)
	mux.HandleFunc("/hostname", HostnameHandler)
	mux.HandleFunc("/headers", HeadersHandler)
	mux.HandleFunc("/", NslookupHandler)

	log.Printf("server is listening at %s", addr)
	http.ListenAndServe(addr, HandleFunc(AllowMethods(LoggingMiddleware(mux), methods), "/headers"))

}

var methods = []string{"/hash", "/hostname", "/headers", "/favicon.ico"}

func AllowMethods(next http.Handler, methods []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, v := range methods {
			if v == r.URL.Path {
				next.ServeHTTP(w, r)
				return
			}
		}
		log.Println("Error", r.URL.Path, "is not allowed")

	})
}

func HandleFunc(next http.Handler, method string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == method {
			next.ServeHTTP(w, r)
		} else {
			log.Println("Error", r.URL.Path, "is not allowed, only", method, " is allowed")
		}

	})
}

func LoggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		defer func() { log.Println(r.URL.Path, " , time: ", time.Since(start)) }()
	})

}

func HostnameHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Fprintf(w, "Hostname: %s", hostname)

}

func HeadersHandler(w http.ResponseWriter, r *http.Request) {
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", name, value)
		}
	}
}

func GetHashForAlgorithm(algorithm string, text string) string {
	var hash string
	switch strings.ToUpper(algorithm) {
	case "MD5":
		hash = fmt.Sprintf("%x", md5.Sum([]byte(string(text))))
	case "SHA256":
		hash = fmt.Sprintf("%x", sha256.Sum256([]byte(text)))
	case "SHA1":
		hash = fmt.Sprintf("%x", sha1.Sum([]byte(text)))
	case "SHA512":
		var sha512Hasher = sha512.New()
		hash = fmt.Sprintf("%x", sha512Hasher.Sum([]byte(text)))
	default:
		hash = ""
	}
	return hash
}

func HashHandler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	text := query.Get("text")
	algorithm := query.Get("algorithm")

	if algorithm == "" {
		algorithm = "undefined"
	}

	hash := GetHashForAlgorithm(algorithm, text)

	hashMap := make(map[string]string)

	hashMap["text"] = text
	hashMap["algorithm"] = algorithm
	hashMap["hash"] = hash

	jData, err := json.Marshal(hashMap)
	if err != nil {
		fmt.Fprint(w, "Marshaling error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func NslookupHandler(w http.ResponseWriter, r *http.Request) {

	reg, _ := regexp.Compile(`^nslookup/((([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9]))$`)
	match := reg.FindStringSubmatch(r.URL.Path[1:])

	if match != nil {
		fmt.Fprintf(w, "Server: %s\n", match[1])
		fmt.Fprintf(w, "\n")

		NSs, err := net.LookupNS(match[1])
		if err != nil {
			fmt.Fprintf(w, "Could not get NS record: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(w, "NS: \n")

		for _, ns := range NSs {
			fmt.Fprintf(w, "%s\n", ns.Host)
		}

		fmt.Fprintf(w, "\n")

		MXs, err := net.LookupMX(match[1])
		if err != nil {
			fmt.Fprintf(w, "Could not get NS record: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(w, "MX: \n")

		for _, mx := range MXs {
			fmt.Fprintf(w, "%s\n", mx.Host)
		}

		fmt.Fprintf(w, "\n")

		ips, err := net.LookupIP(match[1])
		if err != nil {
			fmt.Fprintf(w, "Could not get IPs: %v\n", err)
			os.Exit(1)
		}

		for _, ip := range ips {
			fmt.Fprintf(w, "Address: %s\n", ip.String())
		}
	} else {
		fmt.Fprintf(w, "Unknown parameter: %s!", r.URL.Path[1:])
	}
}

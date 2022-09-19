package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Write here your routes and handlers
	fmt.Println("Running server on http://localhost:5000")
	http.HandleFunc("/hash", HashHandler)
	http.HandleFunc("/hostname", HostnameHandler)
	http.HandleFunc("/headers", HeadersHandler)
	http.HandleFunc("/", NslookupHandler)
	http.ListenAndServe(":5000", nil)
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
		// Create sha-512 hasher
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

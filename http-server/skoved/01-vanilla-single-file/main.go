package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

// routes
const (
	hostNameRoute = "/hostname"
	nsLookupRoute = "/nslookup/"
	hashRoute     = "/hash/"
	headersRoute  = "/headers"
)

// hashing algorithms
const (
	md5Algo    = "md5"
	sha256Algo = "sha256"
	sha512Algo = "sha512"
)

const handleFuncName = "HandleFunc"

var allHandlers RequestHandlers = []requestHandler{
	newHostNameHandler(),
	newNsLookupHandler(),
	newHashHandler(),
	newHeaderHandler(),
}

// interface for requestHandlers. Makes it easier to register handlers when the
// server starts
type requestHandler interface {
	Path() string                                  // path for the endpoint
	HandleFunc(http.ResponseWriter, *http.Request) // the handler function for the endpoint
}

type RequestHandlers []requestHandler

// Calls http.HandleFunc for each path - handler pair
func (rh RequestHandlers) RegisterHandlers() {
	for _, handler := range rh {
		http.HandleFunc(handler.Path(), handler.HandleFunc)
	}
}

type hostNameHandler struct {
	path string
}

// Uses os.Hostname to look up the hostname of the system on when
// http.Request.Method == GET and writes the hostname to
// http.ResponseWriter.
// Sets the Content-Type to "text/plain".
// If http.Request.Method != GET OR cannot obtain the hostname,
// an error message is written to http.ResponseWriter
func (hnh hostNameHandler) HandleFunc(w http.ResponseWriter, r *http.Request) {
	var resp string
	w.Header().Set("Content-Type", "text/plain")
	fmt.Printf("Info: got %s request for %s\n", r.Method, hnh.path)

	if notHttpGet(r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		resp = fmt.Sprintf("HTTP request to %s must use the %s method", hnh.path, http.MethodGet)
		io.WriteString(w, resp)
		return
	}

	resp, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error: unable to lookup hostname\n%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		resp = "Error: Unable to look up the hostname of the server\n"
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Printf("Info: Hostname - %s\n", resp)
	}

	fmt.Println("Info: Writing response")
	io.WriteString(w, resp)
}

func (hnh hostNameHandler) Path() string {
	return hnh.path
}

func newHostNameHandler() *hostNameHandler {
	return &hostNameHandler{path: hostNameRoute}
}

type nsLookupHandler struct {
	path string
}

func (nlh nsLookupHandler) HandleFunc(w http.ResponseWriter, r *http.Request) {
	var resp string
	w.Header().Set("Content-Type", "text-plain")
	fmt.Printf("Info: got %s request for %s\n", r.Method, nlh.path)

	if notHttpGet(r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		resp = fmt.Sprintf("HTTP request to %s must use the %s method", nlh.path, http.MethodGet)
		io.WriteString(w, resp)
		return
	}

	// Get the path param 'hostname'
	hostname := strings.TrimPrefix(r.URL.String(), nsLookupRoute)
	if hostname == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Error: Bad Request. Missing /:hostname path parameter")
		return
	}

	fmt.Printf("Looking up IP addr for hostname %s\n", hostname)

	ips, err := net.LookupIP(hostname)
	if err != nil {
		resp = fmt.Sprintf("Error: unable to look up IP for %s\n", hostname)
		fmt.Print(resp)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		stringIps := make([]string, 0, len(ips))
		for _, ip := range ips {
			stringIps = append(stringIps, ip.String())
		}
		resp = strings.Join(stringIps, " ")
	}

	io.WriteString(w, resp)
}

func (nlh nsLookupHandler) Path() string {
	return nlh.path
}

func newNsLookupHandler() nsLookupHandler {
	return nsLookupHandler{path: nsLookupRoute}
}

type hashHandler struct {
	path string
}

type hashResponse struct {
	Algorithm string `json:"algorithm"`
	Hash      string `json:"hash"`
	Text      string `json:"text"`
}

func (hh hashHandler) HandleFunc(w http.ResponseWriter, r *http.Request) {
	var resp string
	jsonResp := hashResponse{}
	w.Header().Set("Content-Type", "text-plain")
	fmt.Printf("Info: got %s request for %s\n", r.Method, hh.path)

	if notHttpGet(r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		resp = fmt.Sprintf("HTTP request to %s must use the %s method", hh.path, http.MethodGet)
		io.WriteString(w, resp)
		return
	}

	query := r.URL.Query()
	jsonResp.Text = query["text"][0]
	jsonResp.Algorithm = query["algorithm"][0]
	if jsonResp.Text == "" || jsonResp.Algorithm == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("Info: missing either 'text' or 'algorithm' query param\n")
		io.WriteString(w, "Error: query params 'text' and 'algorithm must be provided in http request")
	}

	var hash hash.Hash
	switch jsonResp.Algorithm {
	case md5Algo:
		hash = md5.New()
	case sha256Algo:
		hash = sha256.New()
	case sha512Algo:
		hash = sha512.New()
	default:
		resp = fmt.Sprintf("algorithm %s is not supported\n", jsonResp.Algorithm)
		fmt.Printf("Info: %s", resp)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Error: "+resp)
		return
	}

	io.WriteString(hash, jsonResp.Text)
	jsonResp.Hash = fmt.Sprintf("%x", hash.Sum(nil))

	byteResp, err := json.Marshal(jsonResp)
	if err != nil {
		resp = fmt.Sprintf("Error: Could not marshal json response %+v\n%s\n", jsonResp, err)
		fmt.Print(resp)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, resp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(byteResp)
}

func (hh hashHandler) Path() string {
	return hh.path
}

func newHashHandler() hashHandler {
	return hashHandler{path: hashRoute}
}

type headerHandler struct {
	path string
}

func (hh headerHandler) HandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text-plain")
	fmt.Printf("Info: got %s request for %s\n", r.Method, hh.path)

	if notHttpGet(r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, fmt.Sprintf("HTTP request to %s must use the %s method", hh.path, http.MethodGet))
		return
	}

	io.WriteString(w, fmt.Sprintf("%s", r.Header))
}

func (hh headerHandler) Path() string {
	return hh.path
}

func newHeaderHandler() headerHandler {
	return headerHandler{path: headersRoute}
}

func notHttpGet(httpMethod string) bool {
	return httpMethod != http.MethodGet
}

func main() {
	allHandlers.RegisterHandlers()
	err := http.ListenAndServe(":3000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

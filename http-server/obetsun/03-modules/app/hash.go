package app

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

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

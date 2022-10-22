package app_test

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func getResponse(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(b)

}

func TestNewServer(t *testing.T) {
	hostname, _ := os.Hostname()
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(hostname))
	}))

	defer s.Close()

	if want, got := hostname, getResponse(s.URL); want != got {
		t.Fatalf("expected hostname %s, got %s instead", want, got)
	}
}

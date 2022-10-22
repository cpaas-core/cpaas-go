package app

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandleHostname(t *testing.T) {
	hostname, err := os.Hostname()
	req := httptest.NewRequest(http.MethodGet, "/hostname", nil)
	w := httptest.NewRecorder()

	HandleHostname(w, req)

	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected status %d, got %d instead", want, got)
	}
	if want, got := hostname, string(data); want != got {
		t.Fatalf("expected hostname %s, got %s instead", want, got)
	}
}

package app_test

import (
	"application/app"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewServer(t *testing.T) {
	server := app.NewServer()

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusNotFound {
		t.Fatalf("Wanted 404, got %v", w.Result().StatusCode)
	}
}

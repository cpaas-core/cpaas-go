package app_test

import (
	"application/app"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHostname(t *testing.T) {
	req := httptest.NewRequest("GET", "/hostname", nil)
	w := httptest.NewRecorder()

	app.HandleHostname(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Fatal("Wanted 200")
	}

	if fmt.Sprint(w.Body) != "hemartin.remote.csb" {
		t.Fatalf("Wanted 'hemartin.remote.csb' got '%v'", w.Body)
	}
}

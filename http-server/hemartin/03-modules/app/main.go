package app

import (
	"fmt"
	"net/http"
)

func NewHandler() *http.ServeMux {
	fmt.Println("Returning ServeMux")
	return &http.ServeMux{}
}

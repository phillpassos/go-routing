package main

import (
	"net/http"

	"github.com/phillpassos/go-routing/serve"
)

func main() {
	s := serve.New()
	http.HandleFunc("/", s.ServeHTTP)
	http.ListenAndServe(":8000", nil)
}

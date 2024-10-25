package main

import (
	"fmt"
	"net/http"
)

type Handler struct {
	Name string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "name: %s, url: %v", h.Name, r.URL.String())
}
func main() {
	testHandler := &Handler{Name: "test"}
	http.Handle("/tests/", testHandler)
	rootHandler := &Handler{Name: "root"}
	http.Handle("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}

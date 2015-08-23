package main

import (
	"io"
	"net/http"
	"log"
	"fmt"
	"html"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func rest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
			
func main() {
	http.HandleFunc("/", rest)
	log.Fatal(http.ListenAndServe(":8000", nil))
}


package main

import (
	"net/http"
	"log"
	"fmt"
	"html"
)

func rest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
			
func main() {
	http.HandleFunc("/", rest)
	log.Fatal(http.ListenAndServe(":8000", nil))
}


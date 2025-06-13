package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Request to %s handled by go", r.URL.Path)
	})
	fmt.Println("Server starting")
	http.ListenAndServe(":8080", nil)
}

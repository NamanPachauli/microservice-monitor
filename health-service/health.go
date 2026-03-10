package main

import (
	"fmt"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func main() {
	http.HandleFunc("/health", health)
	fmt.Println("Starting server on port 7001")
	http.ListenAndServe(":7001", nil)
}

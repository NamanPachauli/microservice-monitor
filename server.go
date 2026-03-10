package main

import (
	"fmt"
	"net/http"
)

func main() {
	connectDB()
	go startHealthChecker()

	http.HandleFunc("/services", getServices)
	http.HandleFunc("/register", registerService)

	fmt.Println("Server running on port 5001")

	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		panic(err)
	}
}

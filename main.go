package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServe(":4989", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
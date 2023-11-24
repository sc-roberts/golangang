package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	port := ":8081"

	fmt.Printf("Server listening on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

}

// Graceful exit needed that shuts down the listener/kills the process

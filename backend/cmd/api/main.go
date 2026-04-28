package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server on :8080")

	http.HandleFunc("/health", healthCheck)

	http.ListenAndServe(":8080", nil)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	// Could expand to perform actual health checks. e.g. database connectivity, external service availability, etc.
	log.Println("Performing health check...")
	fmt.Fprintf(w, "It's alive!")
}

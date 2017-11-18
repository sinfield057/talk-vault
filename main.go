package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
		log.Printf("$PORT is not set. Defaulting to %s...\n", port)
	}

	router := http.NewServeMux()
	router.HandleFunc("/", testing)

	http.ListenAndServe(":"+port, router)
}

func testing(w http.ResponseWriter, r *http.Request) {
	log.Println("new hit")
	io.WriteString(w, "Server under construction!")
}

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
		log.Println("$PORT is not set. Defaulting to 8080...")
		port = "8080"
	}

	router := http.NewServeMux()
	router.HandleFunc("/", testing)

	http.ListenAndServe(":"+port, router)
}

func testing(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Server under construction!")
}

package main

import (
	"io"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", testing)

	http.ListenAndServe(":8080", router)
}

func testing(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Server under construction!")
}

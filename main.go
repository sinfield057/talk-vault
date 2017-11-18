package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
		log.Printf("$PORT is not set. Defaulting to %s...\n", port)
	}

	router := httprouter.New()
	//router.ServeFiles("/", http.FileServer(http.Dir("./client/dist")))
	//router.HandleFunc("/*", testing)
	router.GET("/altceva", testing)
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./client/dist/index.html")
	})

	http.Handle("/", http.FileServer(http.Dir("./client/dist/")))
	http.ListenAndServe(":"+port, router)
}

func testing(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "404 from backend")
}

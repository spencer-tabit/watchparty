package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		w.Write([]byte("Welcome to WatchParty!"))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Page not found!"))
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)

	port := ":8080"
	fmt.Println("WatchParty server running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(port, mux))
}

package main

import (
	"html/template"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	default:
		w.WriteHeader(http.StatusNotFound)
		tmpl, err := template.ParseFiles("templates/404.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
		}
		tmpl.Execute(w, nil)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/", homeHandler)

	port := ":8080"
	log.Println("WatchParty server running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(port, mux))
}

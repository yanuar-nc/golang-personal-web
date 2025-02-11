package main

import (
	"fmt"
	"log"
	"net/http"

	"learning/personal/internal/transport/web"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	handler := web.Handler{}
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("assets/statics"))))

	r.HandleFunc("/", handler.Home)
	r.HandleFunc("/contact", handler.Contact)
	r.HandleFunc("/resume", handler.Resume)

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	port := "0.0.0.0:8080"
	log.Default().Println("server is running ", port)
	http.ListenAndServe(port, r)
}

package web

import (
	"fmt"
	"net/http"
	"text/template"
)

type PageData struct {
	Title string
	Name  string
}

type Handler struct{}

func (h Handler) layoutBuilder(page string) []string {
	return []string{
		"internal/transport/web/layouts/layout.html",
		"internal/transport/web/layouts/nav.html",
		"internal/transport/web/layouts/footer.html",

		fmt.Sprintf("internal/transport/web/pages/%s.html", page),
	}
}

func (h Handler) Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		h.layoutBuilder("home")...,
	)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	// vars := mux.Vars(r)
	// title := vars["title"]
	// page := vars["page"]
	data := PageData{
		Title: "Home Page",
		Name:  "John Doe",
	}

	tmpl.Execute(w, data)
	// fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}

func (h Handler) Contact(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		h.layoutBuilder("contact")...,
	)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func (h Handler) Resume(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		h.layoutBuilder("resume")...,
	)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

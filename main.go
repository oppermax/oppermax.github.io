package main

import (
	"html/template"
	"net/http"
	log "github.com/sirupsen/logrus"
)

var tpl *template.Template
func init() {
	tpl = template.Must(tpl.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.Handle("/resources/images/", http.StripPrefix("/resources/images", http.FileServer(http.Dir("./resources/images"))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.WithError(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "home.gohtml", nil)
	if err != nil {
		log.WithError(err)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	if err != nil {
		log.WithError(err)
	}
}
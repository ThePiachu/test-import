package main

import (
	appengine "google.golang.org/appengine/v2"
	"html/template"
	"net/http"

	"github.com/ThePiachu/Go/Log"
	gobcy "github.com/ThePiachu/gobcy/v2"
)

var MainTemplate *template.Template
var BCY gobcy.API

func init() {
	http.HandleFunc("/", hello)

	//Admin functions
	http.HandleFunc("/update", update)

	MainTemplate, _ = template.ParseFiles("html/main.html")
}

func main() {
	appengine.Main()
}

func hello(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	MainTemplate.Execute(w, nil)
}

func update(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	Work.Work(c)
	MainTemplate.Execute(w, nil)
}

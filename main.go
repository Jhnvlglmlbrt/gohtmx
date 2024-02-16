package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello World!")

	handlerData := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "Edge of Tomorrow", Director: "Doug Liman"},
				{Title: "Breaking Bad", Director: "Vince Gilligan"},
				{Title: "Troy", Director: "Wolfgang Petersen"},
			},
		}
		tmpl.Execute(w, films)
	}

	handlerList := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}
	http.HandleFunc("/", handlerData)
	http.HandleFunc("/add-film/", handlerList)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

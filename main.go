package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Product struct {
	Name     string
	Producer string
}

func main() {
	fmt.Println("Hello World!")

	handlerData := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		products := map[string][]Product{
			"Products": {
				{Name: "iPhone 13", Producer: "Apple"},
				{Name: "Galaxy S21", Producer: "Samsung"},
				{Name: "Pixel 6", Producer: "Google"},
			},
		}
		tmpl.Execute(w, products)
	}

	handlerList := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		name := r.PostFormValue("name")
		producer := r.PostFormValue("producer")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "product-list-element", Product{Name: name, Producer: producer})
	}
	http.HandleFunc("/", handlerData)
	http.HandleFunc("/add-product/", handlerList)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

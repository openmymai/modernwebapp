package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	// sum := addValues(2, 2)
	// _, _ = fmt.Fprintf(w, fmt.Sprintf("This is About page and 2 + 2 is %d", sum))
	renderTemplate(w, "about.page.tmpl")
}

// AddValues adds two integers
func addValues(x, y int) int {
	return x + y
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

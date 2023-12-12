package handlers

import (
	"net/http"

	_ "github.com/openmymai/modern_webapp_in_golang/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.renderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	// sum := addValues(2, 2)
	// _, _ = fmt.Fprintf(w, fmt.Sprintf("This is About page and 2 + 2 is %d", sum))
	render.renderTemplate(w, "about.page.tmpl")
}

// AddValues adds two integers
func addValues(x, y int) int {
	return x + y
}

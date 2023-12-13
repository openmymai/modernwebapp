package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/openmymai/modern_webapp_in_golang/pkg/config"
	"github.com/openmymai/modern_webapp_in_golang/pkg/handlers"
	"github.com/openmymai/modern_webapp_in_golang/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	// http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil) // add _ , ignore if error
}

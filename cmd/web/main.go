package main

import (
	"fmt"
	"log"
	"net/http"
	"web_app/pkg/config"
	"web_app/pkg/handlers"
	"web_app/pkg/render"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("server started at port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}

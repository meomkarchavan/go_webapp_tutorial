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
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("server started at port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}

package main

import (
	"fmt"
	"net/http"
	"web_app/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println("server started")
	http.ListenAndServe(portNumber, nil)
}

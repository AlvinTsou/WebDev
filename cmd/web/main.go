package main

import (
	"fmt"
	"log"

	//"html/template"
	"net/http"

	"github.com/AlvinTsou/WebDev/pkg/config"
	"github.com/AlvinTsou/WebDev/pkg/handlers"
	"github.com/AlvinTsou/WebDev/pkg/rander"
)

// const cant be changed
const portNumber = ":8080"

// main application function
func main() {
	// using 2nd templates loading with config.AppConfig
	var app config.AppConfig

	tc, err := rander.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}

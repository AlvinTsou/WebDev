package main

import (
	"fmt"
	"log"

	//"html/template"
	"net/http"

	"github.com/AlvinTsou/WebDev/pkg/config"
	"github.com/AlvinTsou/WebDev/pkg/handlers"
	"github.com/AlvinTsou/WebDev/pkg/render"
)

// const cant be changed
const portNumber = ":8080"

// main application function
func main() {
	// using 2nd templates loading with getring the template cache from config.go
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplatesCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	// 2nd solution

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s\n", portNumber)

	//_ = http.ListenAndServe(portNumber, nil)

	// another way to create a server without using http.HandleFunc
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

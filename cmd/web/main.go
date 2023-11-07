package main

import (
	"fmt"
	//"html/template"
	"net/http"

	"github.com/AlvinTsou/WebDev/pkg/handlers"
)

// const cant be changed
const portNumber = ":8080"

// main application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}

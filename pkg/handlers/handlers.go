package handlers

import (
	//"fmt"
	//"html/template"
	"net/http"
)

// const cant be changed
//const portNumber = ":8080"

// Home is a function that handles the home route
func Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, Alvin's Home!\n")
	RenderTemplate(w, "home.page.tmpl")
}

// About is a function that handles the about route
func About(w http.ResponseWriter, r *http.Request) {
	//sum := AddValues(2, 2)
	//fmt.Fprintf(w, "Alvin's About! And 2+2 is %d\n", sum)
	RenderTemplate(w, "about.page.tmpl")
}

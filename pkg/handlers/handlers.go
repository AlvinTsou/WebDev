package handlers

import (
	"net/http"

	"github.com/AlvinTsou/WebDev/pkg/render"
)

// Home is a function that handles the home route
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is a function that handles the about route
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

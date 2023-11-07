package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders templates using html/template
func RenderTemplateTEST(w http.ResponseWriter, tmpl string) {
	//if use base.layout.tmpl, then need to increase string "./templates/base.layout.tmpl" after +tmpl
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	//parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates:", err)
		return
	}

}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//check to see if we already have the template in the cache
	_, inMap := tc[t]
	if !inMap {
		//need to create the template set and add it to the cache
		log.Println("Create template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// we have the template set already, get it from the cache
		log.Println("Get template from cache")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates:", err)
		return
	}
}

// To create a template cache
func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl",
	}

	//parse the template files
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to the cache (map)
	tc[t] = tmpl
	return nil

}

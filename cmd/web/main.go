package main

import (
	"fmt"
	//"html/template"
	"net/http"
)

// const cant be changed
const portNumber = ":8080"

// Home is a function that handles the home route
/* func Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, Alvin's Home!\n")
	renderTemplate(w, "home.page.tmpl")
} */

// About is a function that handles the about route
/* func About(w http.ResponseWriter, r *http.Request) {
	//sum := AddValues(2, 2)
	//fmt.Fprintf(w, "Alvin's About! And 2+2 is %d\n", sum)
	renderTemplate(w, "about.page.tmpl")
} */

/* func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates:", err)
		return
	}

	//path := fmt.Sprintf("web/templates/%s.html", tmpl)
	//t, err := template.ParseFiles(path)
	//if err != nil {
	//	log.Println("Cannot parse template")
	//	return
	//}
	//err = t.Execute(w, nil)
	//if err != nil {
	//	log.Println("Cannot execute template")
	//	return
	//}
} */

// AddValues adds two integers and return sum, and Add with capital A is exported
/* func AddValues(x, y int) int {
	return x + y
} */

/* func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := DivideValues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero\n")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
}

func DivideValues(x, y float32) (float32, error) {
	if y <= 0.0 {
		err := fmt.Errorf("Cannot divide by zero")
		return 0.0, err
	}
	result := x / y
	return result, nil
} */

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	//http.HandleFunc("/Divide", Divide)

	fmt.Printf("Starting application on port %s\n", portNumber)

	/* 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	   	n, err := fmt.Fprintf(w, "Hello, Alvin!")
	   	if err != nil {
	   		fmt.Println(err)

	   	}
	   	fmt.Printf("Number of ytes written: %d\n", n)
	   	}) */

	_ = http.ListenAndServe(portNumber, nil)
}

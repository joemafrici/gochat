package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// ***********************************************
func main() {
	fmt.Println("hello world")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ***********************************************
func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("in handler")
	log.Println(r.URL.Path)
	switch r.URL.Path {
	case "/login-submit":
		loginSubmit(w, r)
	case "/":
		serveIndex(w, r)
	default:
		serveIndex(w, r)
	}
}

// ***********************************************
func loginSubmit(w http.ResponseWriter, r *http.Request) {
	log.Println("in loginSubmit")
	username := r.FormValue("username")
	password := r.FormValue("password")
	log.Println("Username: ", username, " Password: ", password)
}

// ***********************************************
func serveIndex(w http.ResponseWriter, r *http.Request) {
	var fileName = "index.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		log.Println("Error: Parsing template")
		os.Exit(1)
	}
	t.ExecuteTemplate(w, fileName, nil)
}

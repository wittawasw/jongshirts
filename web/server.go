package server

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type ShirtPageData struct {
    PageTitle string
}

// start the web server r.HandleFunc("/books/{title}", CreateBook).Methods("POST")

func Start() {
	fmt.Println("Starting server")
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8080", r)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/home.html")
	if err != nil {
		fmt.Println(err)
	}
	data := ShirtPageData{
		PageTitle: "Home page",
	}
		
	
	tmpl.Execute(w, data)

}

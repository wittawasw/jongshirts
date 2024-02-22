package server

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type ShirtPageData struct {
	PageTitle string
	ShirtList []ShirtList
}

type ShirtList struct {
	Name  string
	Size  string
	Price string
	Color string
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
		ShirtList: []ShirtList{
			{Name: "shirt 1", Price: "100", Color: "Red", Size: "XL"},
			{Name: "shirt 2", Price: "50", Color: "Green", Size: "L"},
			{Name: "shirt 3", Price: "300", Color: "Blue-green", Size: "S"},
			{Name: "shirt 4", Price: "77", Color: "Black", Size: "XXXL"},
		},
	}

	tmpl.Execute(w, data)

}

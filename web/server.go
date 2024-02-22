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
	Id    int
	Name  string
	Size  string
	Price string
	Color string
}

type Incart struct {
	Name []string
}

// start the web server r.HandleFunc("/books/{title}", CreateBook).Methods("POST")

func Start() {
	fmt.Println("Starting server")
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/cart", cartHandler)
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
			{Id: 1, Name: "shirt 1", Price: "100", Color: "Red", Size: "XL"},
			{Id: 2, Name: "shirt 2", Price: "50", Color: "Green", Size: "L"},
			{Id: 3, Name: "shirt 3", Price: "300", Color: "Blue-green", Size: "S"},
			{Id: 4, Name: "shirt 4", Price: "77", Color: "Black", Size: "XXXL"},
		},
	}

	tmpl.Execute(w, data)

}

func cartHandler(w http.ResponseWriter, r *http.Request) {

	var SelectedShirts []string

	r.ParseForm()
	for key, _ := range r.Form {
		SelectedShirts = append(SelectedShirts, key)
	}

	fmt.Println(SelectedShirts)

	tmpl, err := template.ParseFiles("web/templates/detail.html")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(w, SelectedShirts)

}

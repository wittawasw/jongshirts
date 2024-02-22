package server

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/oddsteam/jongshirts/internal/db"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
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
	r.HandleFunc("/cart", cartHandler).Methods("POST")
	r.HandleFunc("/showcart", ShowCart)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/auth", authenticationHandler)
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
	client := db.NewClient()
	// var SelectedShirts []string

	r.ParseForm()
	// ctx := context.Background()

	for key, _ := range r.Form {
		client.LPush("selectedShirt", key)
	}

	http.Redirect(w, r, "/showcart", http.StatusSeeOther)
}

func ShowCart(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/detail.html")
	if err != nil {
		fmt.Println(err)
	}

	client := db.NewClient()
	data, err := client.LRange("selectedShirt", 0, -1).Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)

	tmpl.Execute(w, data)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/login.html")
	if err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(w, nil)

}

func authenticationHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "sessions")
	// Authentication goes here
	// ...
	email:=   r.FormValue("email")
	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Values["username"] = email
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)

	
}

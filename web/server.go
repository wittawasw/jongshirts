package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// start the web server r.HandleFunc("/books/{title}", CreateBook).Methods("POST")

func Start() {
	fmt.Println("Starting server")
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8080", r)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is Home handler")

}

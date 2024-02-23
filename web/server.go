package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oddsteam/jongshirts/web/handlers"
)

func Start() {
	fmt.Println("Starting server")
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/cart", handlers.CartHandler).Methods("POST")
	r.HandleFunc("/showcart", handlers.ShowCartHandler)
	r.HandleFunc("/login", handlers.LoginHandler)
	r.HandleFunc("/auth", handlers.AuthenticationHandler)

	http.ListenAndServe(":8080", r)
}

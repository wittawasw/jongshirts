package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// start the web server

func Start() {
	fmt.Println("Starting server")
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello World")
	})
	http.ListenAndServe(":8080", r)

}

package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"

	"github.com/oddsteam/jongshirts/internal/sessions"
	"github.com/oddsteam/jongshirts/web/handlers"
)

func Start() {
	fmt.Println("Starting server")

	gothic.Store = sessions.Store

	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	callbackUrl := "http://" + host + ":" + port + "/oauth/google/callback"

	goth.UseProviders(
		google.New(googleClientId, googleClientSecret, callbackUrl, "email"),
	)

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/cart", handlers.CartHandler).Methods("POST")
	r.HandleFunc("/showcart", handlers.ShowCartHandler)
	r.HandleFunc("/login", handlers.LoginHandler)
	r.HandleFunc("/logout", handlers.LogOutHandler)

	r.HandleFunc("/auth", handlers.AuthenticationHandler)

	r.HandleFunc("/oauth/{provider}/callback", handlers.OAuthCallbackHandler)
	r.HandleFunc("/oauth/{provider}", handlers.OAuthHandler)

	http.ListenAndServe(":"+port, r)
}

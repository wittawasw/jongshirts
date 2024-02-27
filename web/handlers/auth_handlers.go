package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/oddsteam/jongshirts/internal/sessions"
)

func AuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.InitSession(r)

	email := r.FormValue("email")

	session.Values["authenticated"] = true
	session.Values["username"] = email
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func OAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	session, _ := sessions.InitSession(r)

	session.Values["authenticated"] = true
	session.Values["username"] = user.Email
	session.Save(r, w)

	t, _ := template.ParseFiles("web/templates/success.html")

	t.Execute(w, user)
}

func OAuthHandler(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

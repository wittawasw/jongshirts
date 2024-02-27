package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/oddsteam/jongshirts/internal/sessions"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.InitSession(r)

	if session.Values["authenticated"] == true {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if session.Values["username"] != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("web/templates/login.html")
	if err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(w, nil)
}

func LogOutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.InitSession(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Clear session data
	for key := range session.Values {
		delete(session.Values, key)
	}

	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

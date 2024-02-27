package sessions

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte(os.Getenv("SECRET_KEY"))
	Store = sessions.NewCookieStore(key)
)

func InitSession(r *http.Request) (*sessions.Session, error) {
	session, err := Store.Get(r, "sessions")
	if session.IsNew { //Set some cookie options
		session.Options.Domain = os.Getenv("HOST")
		session.Options.MaxAge = 0
		session.Options.HttpOnly = false
		session.Options.Secure = true
	}
	return session, err
}

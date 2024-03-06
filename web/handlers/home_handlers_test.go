package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oddsteam/jongshirts/internal/sessions"
	"github.com/oddsteam/jongshirts/web/handlers"
)

func TestHomeHandler(t *testing.T) {
	handlers.SetTemplateDir("../templates/")

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	session, _ := sessions.InitSession(req)
	session.Values["username"] = "testuser"
	session.Save(req, httptest.NewRecorder())

	w := httptest.NewRecorder()
	handlers.HomeHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, resp.StatusCode)
	}
}

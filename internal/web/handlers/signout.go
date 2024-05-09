package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

func Signout(logger *log.Logger, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "user-session")
		session.Values["authenticated"] = false
		session.Save(r, w)
		w.Header().Set("HX-Redirect", "/signin")
		w.WriteHeader(http.StatusOK)

	}
}

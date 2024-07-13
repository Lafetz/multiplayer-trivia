package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/sessions"
)

func Signout(logger *slog.Logger, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "user-session")
		session.Values["authenticated"] = false
		session.Save(r, w)
		w.Header().Set("HX-Redirect", "/signin")
		w.WriteHeader(http.StatusOK)

	}
}

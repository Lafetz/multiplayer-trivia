package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/sessions"
)

func Signout(logger *slog.Logger, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "user-session")
		if err != nil {
			ServerError(w, r, err, logger)
		}
		session.Values["authenticated"] = false
		err = session.Save(r, w)
		if err != nil {

			ServerError(w, r, err, logger)
		}
		w.Header().Set("HX-Redirect", "/signin")
		w.WriteHeader(http.StatusOK)

	}
}

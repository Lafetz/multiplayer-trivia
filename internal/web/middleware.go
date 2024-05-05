package web

import (
	"context"
	"net/http"
)

func (a *App) requireAuth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve session
		session, err := a.store.Get(r, "user-session")
		if err != nil || session.Values["authenticated"] != true {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}

		username, ok := session.Values["username"].(string)
		if !ok {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), "username", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

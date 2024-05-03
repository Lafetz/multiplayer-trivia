package web

import (
	"context"
	"fmt"
	"net/http"
)

func (a *App) requireAuth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve session
		session, err := a.store.Get(r, "user-session")
		if err != nil || session.Values["authenticated"] != true {
			fmt.Printf("so %s", session.Values)
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}

		username, ok := session.Values["username"].(string)
		if !ok {
			http.Error(w, "Username not found in session", http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), "username", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

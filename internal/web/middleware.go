package web

import (
	"context"
	"net/http"

	"github.com/Lafetz/showdown-trivia-game/internal/web/handlers"
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
func (app *App) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if e, ok := err.(error); ok {
					handlers.ServerError(w, r, e, app.logger)
				} else {
					app.logger.Println(err)
				}

			}
		}()
		next.ServeHTTP(w, r)
	})
}

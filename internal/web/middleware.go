package web

import (
	"context"
	"net/http"
	"time"

	"github.com/Lafetz/showdown-trivia-game/internal/web/handlers"
	"github.com/prometheus/client_golang/prometheus"
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
		ctx := context.WithValue(r.Context(), handlers.UsernameKey, username)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
func (app *App) recoverPanic(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if e, ok := err.(error); ok {
					handlers.ServerError(w, r, e, app.logger)
				} else {
					app.logger.Error(e.Error())
				}

			}
		}()
		next.ServeHTTP(w, r)
	}
}
func (app *App) requestDuration(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r)
		app.m.ReqDuration.With(prometheus.Labels{"method": r.Method}).Observe(time.Since(now).Seconds())
	}
}

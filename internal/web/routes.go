package web

import (
	"net/http"

	"github.com/Lafetz/showdown-trivia-game/internal/web/form"
	layout "github.com/Lafetz/showdown-trivia-game/internal/web/views/layouts"
	"github.com/Lafetz/showdown-trivia-game/internal/web/views/pages"
)

func (a *App) initAppRoutes() {
	a.router.HandleFunc("GET /signup", func(w http.ResponseWriter, r *http.Request) {
		p := pages.Signup(form.SignupUser{})
		err := layout.Base("Sign up", p).Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	})
	a.router.HandleFunc("POST /signupp", func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 4096)
		err := r.ParseForm()

		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			println("problem parsing form")
			return
		}
		form := form.SignupUser{
			Username: r.PostForm.Get("username"),
			Email:    r.PostForm.Get("email"),
			Password: r.PostForm.Get("password"),
		}
		if !form.Valid() {
			p := pages.Signup(form)
			w.WriteHeader(422)
			err = layout.Base("Sign up", p).Render(r.Context(), w)
			if err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}
			return
		}

	})
}

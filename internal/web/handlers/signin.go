package handlers

import (
	"net/http"

	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
	render "github.com/Lafetz/showdown-trivia-game/internal/web/Render"
	"github.com/Lafetz/showdown-trivia-game/internal/web/form"
	layout "github.com/Lafetz/showdown-trivia-game/internal/web/views/layouts"
	"github.com/Lafetz/showdown-trivia-game/internal/web/views/pages"
	"github.com/gorilla/sessions"
)

func SigninGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := pages.Signin(form.SigninUser{}, "")
		err := layout.Base("Sign in", p).Render(r.Context(), w)
		if err != nil {
			render.InternalServer(w, r)
			return
		}
	}
}
func SigninPost(userservice user.UserServiceApi, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 4096)
		err := r.ParseForm()
		if err != nil {
			render.InternalServer(w, r)
			return
		}
		form := form.SigninUser{
			Email:    r.PostForm.Get("email"),
			Password: r.PostForm.Get("password"),
		}
		if !form.Valid() {
			render.InvalidForm(w, r, form)
			return
		}
		user, err := userservice.GetUser(form.Email)
		if err != nil {
			render.IncorrectPasswordEmail(w, r)
			return
		}

		err = matchPassword(form.Password, user.Password)
		if err != nil {
			if err == ErrInvalidPassword {
				render.IncorrectPasswordEmail(w, r)
			}
			return
		}
		session, _ := store.Get(r, "user-session")
		session.Values["authenticated"] = true
		session.Values["username"] = user.Username
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("HX-Redirect", "/home")
		w.WriteHeader(http.StatusOK)
	}
}

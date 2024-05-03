package handlers

import (
	"net/http"

	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
	render "github.com/Lafetz/showdown-trivia-game/internal/web/Render"
	"github.com/Lafetz/showdown-trivia-game/internal/web/form"
	layout "github.com/Lafetz/showdown-trivia-game/internal/web/views/layouts"
	"github.com/Lafetz/showdown-trivia-game/internal/web/views/pages"
)

func SignupGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := pages.Signup(form.SignupUser{})
		err := layout.Base("Sign up", p).Render(r.Context(), w)
		if err != nil {
			render.InternalServer(w, r)
		}
	}
}
func SignupPost(userService user.UserServiceApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 4096)
		err := r.ParseForm()

		if err != nil {
			render.InternalServer(w, r)
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
				render.InternalServer(w, r)
				return
			}
			return
		}
		hashedPassword, err := hashPassword(form.Password)
		if err != nil {

			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
		user := user.NewUser(form.Username, form.Email, hashedPassword)
		_, err = userService.AddUser(user)
		if err != nil {
			println(err)
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
		p := pages.SignupSuccess()
		w.WriteHeader(201)
		err = layout.Base("Sign up", p).Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	}
}
func Ssig() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
	render "github.com/Lafetz/showdown-trivia-game/internal/web/Render"
	"github.com/Lafetz/showdown-trivia-game/internal/web/form"
	layout "github.com/Lafetz/showdown-trivia-game/internal/web/views/layouts"
	"github.com/Lafetz/showdown-trivia-game/internal/web/views/pages"
)

func SignupGet(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := pages.Signup(form.SignupUser{})
		err := layout.Base("Sign up", p).Render(r.Context(), w)
		if err != nil {
			ServerError(w, r, err, logger)
		}
	}
}
func SignupPost(userService user.UserServiceApi, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.Body = http.MaxBytesReader(w, r.Body, 4096)
		err := r.ParseForm()

		if err != nil {
			ServerError(w, r, err, logger)
			return
		}
		form := form.SignupUser{
			Username: r.PostForm.Get("username"),
			Email:    r.PostForm.Get("email"),
			Password: r.PostForm.Get("password"),
		}
		if !form.Valid() {
			err := render.InvalidFormSignup(w, r, form)
			if err != nil {
				ServerError(w, r, err, logger)
				return
			}
			return
		}
		hashedPassword, err := hashPassword(form.Password)
		if err != nil {
			ServerError(w, r, err, logger)
			return
		}
		u := user.NewUser(form.Username, form.Email, hashedPassword)
		_, err = userService.AddUser(u)
		if err != nil {
			//||
			if errors.Is(err, user.ErrEmailUnique) {
				form.AddError("email", err.Error())
				err := render.InvalidFormSignup(w, r, form)
				if err != nil {
					ServerError(w, r, err, logger)
				}
				return
			}
			if errors.Is(err, user.ErrUsernameUnique) {
				form.AddError("username", err.Error())
				err := render.InvalidFormSignup(w, r, form)
				if err != nil {
					ServerError(w, r, err, logger)
				}
				return
			}
			ServerError(w, r, err, logger)
			return
		}
		p := pages.SignupSuccess()
		w.WriteHeader(201)
		err = layout.Base("Sign up", p).Render(r.Context(), w)
		if err != nil {
			ServerError(w, r, err, logger)
			return
		}
	}
}

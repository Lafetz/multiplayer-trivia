package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
	render "github.com/Lafetz/showdown-trivia-game/internal/web/Render"
	"github.com/Lafetz/showdown-trivia-game/internal/web/form"
	layout "github.com/Lafetz/showdown-trivia-game/internal/web/views/layouts"
	"github.com/Lafetz/showdown-trivia-game/internal/web/views/pages"
	"github.com/gorilla/sessions"
)

const IncorrectEmailPass = "Incorrect Email or Password"

func SigninGet(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := pages.Signin(form.SigninUser{}, "", "")
		err := layout.Base("Sign in", p).Render(r.Context(), w)
		if err != nil {
			ServerError(w, r, err, logger)
			return
		}
	}
}
func SigninPost(userservice user.UserServiceApi, store *sessions.CookieStore, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.Body = http.MaxBytesReader(w, r.Body, 4096)
		err := r.ParseForm()
		if err != nil {
			ServerError(w, r, err, logger)
			return
		}
		form := form.SigninUser{
			Email:    r.PostForm.Get("email"),
			Password: r.PostForm.Get("password"),
		}
		if !form.Valid() {
			if err := render.InvalidFormSignin(w, r, form, "", http.StatusUnprocessableEntity); err != nil {
				ServerError(w, r, err, logger)
			}
			return
		}
		userData, err := userservice.GetUser(form.Email)
		if err != nil {
			if errors.Is(err, user.ErrUserNotFound) {
				if err := render.InvalidFormSignin(w, r, form, IncorrectEmailPass, http.StatusUnauthorized); err != nil {
					ServerError(w, r, err, logger)
				}
				return
			} else {
				ServerError(w, r, err, logger)
			}

		}

		err = matchPassword(form.Password, userData.Password)
		if err != nil {
			if errors.Is(err, ErrInvalidPassword) {
				if err := render.InvalidFormSignin(w, r, form, IncorrectEmailPass, http.StatusUnauthorized); err != nil {
					ServerError(w, r, err, logger)
				}
				return
			}
			ServerError(w, r, err, logger)
			return
		}
		session, _ := store.Get(r, "user-session")
		session.Values["authenticated"] = true
		session.Values["username"] = userData.Username
		err = session.Save(r, w)
		if err != nil {
			ServerError(w, r, err, logger)
			return
		}
		w.Header().Set("HX-Redirect", "/home")
		w.WriteHeader(http.StatusOK)
	}
}

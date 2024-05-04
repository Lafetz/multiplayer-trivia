package render

import (
	"net/http"

	"github.com/Lafetz/showdown-trivia-game/internal/web/entity"
	"github.com/Lafetz/showdown-trivia-game/internal/web/form"
	"github.com/Lafetz/showdown-trivia-game/internal/web/views/components"
	layout "github.com/Lafetz/showdown-trivia-game/internal/web/views/layouts"
	"github.com/Lafetz/showdown-trivia-game/internal/web/views/pages"
)

func Home(w http.ResponseWriter, r *http.Request, username string) error {
	p := pages.Home(username)
	err := layout.Base("Home", p).Render(r.Context(), w)
	if err != nil {
		return err
	}
	return nil

}
func InvalidFormSignin(w http.ResponseWriter, r *http.Request, form form.SigninUser) error {
	w.WriteHeader(http.StatusUnprocessableEntity)
	p := pages.Signin(form, "")
	err := p.Render(r.Context(), w)
	if err != nil {
		return err
	}
	return nil
}
func IncorrectPasswordEmail(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusUnauthorized)
	p := pages.InvalidAuth()
	err := p.Render(r.Context(), w)
	if err != nil {
		return err
	}
	return nil
}
func InvalidFormSignup(w http.ResponseWriter, r *http.Request, form form.SignupUser) error {
	p := pages.Signup(form)
	w.WriteHeader(http.StatusUnprocessableEntity)
	err := layout.Base("Sign up", p).Render(r.Context(), w)
	return err
}
func ActiveGames(w http.ResponseWriter, r *http.Request, rooms []entity.RoomData) error {
	p := components.ActiveGames(rooms)
	err := p.Render(r.Context(), w)
	if err != nil {
		return err
	}
	return nil

}

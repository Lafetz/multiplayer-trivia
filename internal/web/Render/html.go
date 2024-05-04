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
func InvalidForm(w http.ResponseWriter, r *http.Request, form form.SigninUser) error {
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
func InvalidForxm(w http.ResponseWriter, r *http.Request) {}
func ActiveGames(w http.ResponseWriter, r *http.Request, rooms []entity.RoomData) error {
	p := components.ActiveGames(rooms)
	err := p.Render(r.Context(), w)
	if err != nil {
		return err
	}
	return nil

}

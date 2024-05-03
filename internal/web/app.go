package web

import (
	"net/http"

	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
	"github.com/Lafetz/showdown-trivia-game/internal/ws"
	"github.com/gorilla/sessions"
)

type App struct {
	port        int
	router      *http.ServeMux
	userService user.UserServiceApi
	hub         *ws.Hub
	store       *sessions.CookieStore
}

func NewApp(userService user.UserServiceApi, hub *ws.Hub, store *sessions.CookieStore) *App {

	a := &App{
		router:      http.NewServeMux(),
		port:        8080,
		userService: userService,
		hub:         hub,
		store:       store,
	}

	a.initAppRoutes()

	return a
}
func (a *App) Run() error {
	return http.ListenAndServe(":8080", a.router)
}

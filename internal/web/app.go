package web

import (
	"log"
	"net/http"
	"os"

	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
	"github.com/Lafetz/showdown-trivia-game/internal/web/ws"
	"github.com/gorilla/sessions"
)

type App struct {
	port        int
	router      *http.ServeMux
	userService user.UserServiceApi
	hub         *ws.Hub
	store       *sessions.CookieStore
	logger      *log.Logger
}

func NewApp(userService user.UserServiceApi, hub *ws.Hub, store *sessions.CookieStore) *App {

	a := &App{
		router:      http.NewServeMux(),
		port:        8080,
		userService: userService,
		hub:         hub,
		store:       store,
		logger:      log.New(os.Stdout, "", log.LstdFlags),
	}

	a.initAppRoutes()

	return a
}
func (a *App) Run() error {
	return http.ListenAndServe(":8080", a.router)
}

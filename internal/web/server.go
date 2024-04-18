package web

import (
	"net/http"

	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
)

type App struct {
	port        int
	router      *http.ServeMux
	userService user.UserServiceApi
}

func NewApp(userService user.UserServiceApi) *App {

	a := &App{
		router:      http.NewServeMux(),
		port:        8080,
		userService: userService,
	}

	a.initAppRoutes()

	return a
}
func (a *App) Run() error {
	return http.ListenAndServe(":8080", a.router)
}

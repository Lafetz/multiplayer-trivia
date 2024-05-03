package web

import (
	"net/http"

	"github.com/Lafetz/showdown-trivia-game/internal/web/handlers"
)

const (
	wsUrl = "connect:ws://localhost:8080"
)

func (a *App) initAppRoutes() {
	fileServer := http.FileServer(http.Dir("./internal/web/static"))
	s := http.Handler(handlers.Home())
	a.router.Handle("/static/", http.StripPrefix("/static/", fileServer))
	//

	a.router.HandleFunc("GET /signin", handlers.SigninGet())
	a.router.HandleFunc("POST /signin", handlers.SigninPost(a.userService, a.store))
	//

	a.router.HandleFunc("GET /signup", handlers.SignupGet())
	a.router.HandleFunc("POST /signup", handlers.SignupPost(a.userService))
	//

	a.router.HandleFunc("GET /home", a.requireAuth(s))
	a.router.HandleFunc("/create", a.requireAuth(handlers.CreateGet()))
	a.router.HandleFunc("/join/{id}", handlers.Join())
	//ws
	a.router.HandleFunc("/activegames", a.requireAuth(handlers.ActiveGames(a.hub)))
	//
	a.router.HandleFunc("/wscreate", a.requireAuth(handlers.CreateWs(a.hub)))
	a.router.HandleFunc("/wsjoin/{id}", a.requireAuth(handlers.JoinWs(a.hub)))

}

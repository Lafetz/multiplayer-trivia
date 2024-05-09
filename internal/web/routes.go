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
	s := http.Handler(handlers.Home(a.logger))
	a.router.Handle("/static/", http.StripPrefix("/static/", fileServer))
	//

	a.router.HandleFunc("GET /signin", handlers.SigninGet(a.logger))
	a.router.HandleFunc("POST /signin", handlers.SigninPost(a.userService, a.store, a.logger))
	//

	a.router.HandleFunc("GET /signup", handlers.SignupGet(a.logger))
	a.router.HandleFunc("POST /signup", handlers.SignupPost(a.userService, a.logger))
	//

	a.router.HandleFunc("GET /home", a.requireAuth(s))
	a.router.HandleFunc("GET /create", a.requireAuth(handlers.CreateFormGet(a.logger, a.questionService)))
	a.router.HandleFunc("POST /create", a.requireAuth(handlers.CreateFormPost(a.logger, a.questionService)))
	a.router.HandleFunc("GET /join/{id}", handlers.Join(a.logger))
	//ws
	a.router.HandleFunc("/activegames", a.requireAuth(handlers.ActiveGames(a.hub, a.logger)))
	//
	a.router.HandleFunc("/wscreate", a.requireAuth(handlers.CreateWs(a.hub, a.questionService)))
	a.router.HandleFunc("/wsjoin/{id}", a.requireAuth(handlers.JoinWs(a.hub)))

}

package web

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/Lafetz/showdown-trivia-game/internal/web/handlers"
)

//go:embed static/*
var staticFiles embed.FS

func (a *App) initAppRoutes() {
	staticFs, err := fs.Sub(staticFiles, "static")
	if err != nil {
		a.logger.Fatal(err)
	}

	a.router.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFs))))
	a.router.HandleFunc("GET /signin", handlers.SigninGet(a.logger))
	a.router.HandleFunc("POST /signin", handlers.SigninPost(a.userService, a.store, a.logger))
	//

	a.router.HandleFunc("GET /signup", handlers.SignupGet(a.logger))
	a.router.HandleFunc("POST /signup", handlers.SignupPost(a.userService, a.logger))
	a.router.HandleFunc("POST /signout", a.requireAuth(handlers.Signout(a.logger, a.store)))
	//

	a.router.HandleFunc("GET /home", a.requireAuth(handlers.Home(a.logger)))
	a.router.HandleFunc("GET /create", a.requireAuth(handlers.CreateFormGet(a.logger, a.questionService)))
	a.router.HandleFunc("POST /create", a.requireAuth(handlers.CreateFormPost(a.logger, a.questionService)))
	a.router.HandleFunc("GET /join/{id}", handlers.Join(a.logger))
	//ws
	a.router.HandleFunc("/activegames", a.requireAuth(handlers.ActiveGames(a.hub, a.logger)))
	//
	a.router.HandleFunc("/wscreate", a.requireAuth(handlers.CreateWs(a.hub, a.questionService)))
	a.router.HandleFunc("/wsjoin/{id}", a.requireAuth(handlers.JoinWs(a.hub)))

}

package web

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/Lafetz/showdown-trivia-game/internal/web/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//go:embed static/*
var staticFiles embed.FS

func (a *App) initAppRoutes() {
	staticFs, err := fs.Sub(staticFiles, "static")
	if err != nil {
		a.logger.Error(err.Error())
	}
	//

	a.router.Handle("/metrics", promhttp.HandlerFor(a.reg, promhttp.HandlerOpts{}))
	//
	a.router.Handle("GET /static/", a.recoverPanic(http.StripPrefix("/static/", http.FileServer(http.FS(staticFs)))))
	a.router.HandleFunc("GET /signin", a.recoverPanic(handlers.SigninGet(a.logger)))
	a.router.HandleFunc("POST /signin", a.recoverPanic(handlers.SigninPost(a.userService, a.store, a.logger)))
	//

	a.router.HandleFunc("GET /signup", a.recoverPanic(handlers.SignupGet(a.logger)))
	a.router.HandleFunc("POST /signup", a.recoverPanic(handlers.SignupPost(a.userService, a.logger)))
	a.router.HandleFunc("POST /signout", a.recoverPanic(a.requireAuth(handlers.Signout(a.logger, a.store))))
	//

	a.router.HandleFunc("GET /home", a.recoverPanic(a.requireAuth(handlers.Home(a.logger))))
	a.router.HandleFunc("GET /create", a.recoverPanic(a.requestDuration(a.requireAuth(handlers.CreateFormGet(a.logger, a.questionService)))))
	a.router.HandleFunc("POST /create", a.recoverPanic(a.requestDuration(a.requireAuth(handlers.CreateFormPost(a.logger, a.questionService)))))
	a.router.HandleFunc("GET /join/{id}", a.recoverPanic(handlers.Join(a.logger)))
	//ws
	a.router.HandleFunc("/activegames", a.recoverPanic(a.requireAuth(handlers.ActiveGames(a.hub, a.logger))))
	//
	a.router.HandleFunc("/wscreate", a.recoverPanic(a.requireAuth(handlers.CreateWs(a.hub, a.questionService))))
	a.router.HandleFunc("/wsjoin/{id}", a.recoverPanic(a.requireAuth(handlers.JoinWs(a.hub))))

}

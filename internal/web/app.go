package web

import (
	"log"
	"net/http"
	"os"

	"github.com/Lafetz/showdown-trivia-game/internal/core/question"
	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
	"github.com/Lafetz/showdown-trivia-game/internal/web/ws"
	"github.com/gorilla/sessions"
)

type App struct {
	port            int
	router          *http.ServeMux
	userService     user.UserServiceApi
	questionService question.QuestionServiceApi
	hub             *ws.Hub
	store           *sessions.CookieStore
	logger          *log.Logger
}

func NewApp(port int, userService user.UserServiceApi, store *sessions.CookieStore, questionService question.QuestionServiceApi) *App {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	a := &App{
		router:          http.NewServeMux(),
		port:            port,
		userService:     userService,
		hub:             ws.NewHub(logger),
		store:           store,
		questionService: questionService,
		logger:          logger,
	}

	a.initAppRoutes()

	return a
}
func (a *App) Run() error {
	return http.ListenAndServe(":8080", a.router)
}

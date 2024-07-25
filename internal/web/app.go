package web

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Lafetz/showdown-trivia-game/internal/core/question"
	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
	"github.com/Lafetz/showdown-trivia-game/internal/web/metrics"
	"github.com/Lafetz/showdown-trivia-game/internal/web/ws"
	"github.com/gorilla/sessions"
	"github.com/prometheus/client_golang/prometheus"
)

type App struct {
	port            int
	router          *http.ServeMux
	userService     user.UserServiceApi
	questionService question.QuestionServiceApi
	hub             *ws.Hub
	store           *sessions.CookieStore
	logger          *slog.Logger
	m               *metrics.Metrics
	reg             *prometheus.Registry
}

func NewApp(port int, logger *slog.Logger, userService user.UserServiceApi, store *sessions.CookieStore, questionService question.QuestionServiceApi, reg *prometheus.Registry) *App {

	m := metrics.NewMetrics(reg)

	a := &App{
		router:          http.NewServeMux(),
		port:            port,
		userService:     userService,
		hub:             ws.NewHub(logger, m),
		store:           store,
		questionService: questionService,
		logger:          logger,
		reg:             reg,
		m:               m,
	}

	a.initAppRoutes()

	return a
}
func (a *App) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(a.port)), a.router)
}

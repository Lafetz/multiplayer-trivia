package web

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

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

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", strconv.Itoa(a.port)),
		Handler:      a.router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		<-quit

		a.logger.Info("shutting down server")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		shutdownError <- srv.Shutdown(ctx)
	}()
	err := srv.ListenAndServe()

	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}
	a.logger.Info("server stopped")
	return nil
}

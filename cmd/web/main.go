package main

import (
	"log"

	"github.com/Lafetz/showdown-trivia-game/internal/config"
	"github.com/Lafetz/showdown-trivia-game/internal/core/question"
	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
	"github.com/Lafetz/showdown-trivia-game/internal/logger"
	"github.com/Lafetz/showdown-trivia-game/internal/repository"
	triviaapi "github.com/Lafetz/showdown-trivia-game/internal/trivia_api"
	"github.com/Lafetz/showdown-trivia-game/internal/web"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}
	logger := logger.NewLogger(cfg.LogLevel, cfg.Env)
	if err != nil {
		log.Fatal(err)
	}

	store := sessions.NewCookieStore(securecookie.GenerateRandomKey(16), securecookie.GenerateRandomKey(16))
	db, close, err := repository.NewDb(cfg.DbUrl, logger)
	defer close()

	if err != nil {
		log.Fatal(err)
	}
	logger.Info("db connected")
	repo, err := repository.NewStore(db)
	if err != nil {
		log.Fatal(err)
	}
	userservice := user.NewUserService(repo)
	triviaClient := triviaapi.NewTriviaClient()
	questionService := question.NewQuestionService(triviaClient)
	reg := prometheus.NewRegistry()
	app := web.NewApp(cfg.Port, logger, userservice, store, questionService, reg)
	err = app.Run()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

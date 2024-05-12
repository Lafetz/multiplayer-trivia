package main

import (
	"log"

	"github.com/Lafetz/showdown-trivia-game/internal/config"
	"github.com/Lafetz/showdown-trivia-game/internal/core/question"
	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
	"github.com/Lafetz/showdown-trivia-game/internal/repository"
	triviaapi "github.com/Lafetz/showdown-trivia-game/internal/trivia_api"
	"github.com/Lafetz/showdown-trivia-game/internal/web"
	"github.com/gorilla/sessions"
)

func main() {
	cfg := config.NewConfig()
	store := sessions.NewCookieStore([]byte(cfg.HashKey), []byte(cfg.BlockKey))
	db := repository.NewDb(cfg.DbUrl)
	repo := repository.NewStore(db)
	userservice := user.NewUserService(repo)
	triviaClient := triviaapi.NewTriviaClient()
	questionService := question.NewQuestionService(triviaClient)
	app := web.NewApp(cfg.Port, userservice, store, questionService)
	err := app.Run()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

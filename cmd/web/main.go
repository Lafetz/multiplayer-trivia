package main

import (
	"log"
	"os"

	"github.com/Lafetz/showdown-trivia-game/internal/config"
	"github.com/Lafetz/showdown-trivia-game/internal/core/question"
	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
	"github.com/Lafetz/showdown-trivia-game/internal/repository"
	triviaapi "github.com/Lafetz/showdown-trivia-game/internal/trivia_api"
	"github.com/Lafetz/showdown-trivia-game/internal/web"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

func main() {
	cfg := config.NewConfig()
	store := sessions.NewCookieStore(securecookie.GenerateRandomKey(16), securecookie.GenerateRandomKey(16))
	db := repository.NewDb(cfg.DbUrl)
	repo := repository.NewStore(db)
	userservice := user.NewUserService(repo)
	triviaClient := triviaapi.NewTriviaClient()
	questionService := question.NewQuestionService(triviaClient)
	logger := log.New(os.Stdout, "", log.LstdFlags)
	app := web.NewApp(cfg.Port, logger, userservice, store, questionService)
	err := app.Run()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

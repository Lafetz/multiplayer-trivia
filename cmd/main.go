package main

import (
	"log"

	"github.com/Lafetz/showdown-trivia-game/internal/core/question"
	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
	"github.com/Lafetz/showdown-trivia-game/internal/repository"
	triviaapi "github.com/Lafetz/showdown-trivia-game/internal/trivia_api"
	"github.com/Lafetz/showdown-trivia-game/internal/web"
	"github.com/gorilla/sessions"
)

func main() {
	hashKey := "your-generated-hash-key"
	blockKey := "your-generated-block-key"
	store := sessions.NewCookieStore([]byte(hashKey), []byte(blockKey))
	repo := repository.NewStore()
	userservice := user.NewUserService(repo)
	triviaClient := triviaapi.NewTriviaClient()
	questionService := question.NewQuestionService(triviaClient)
	app := web.NewApp(userservice, store, questionService)
	err := app.Run()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

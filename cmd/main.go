package main

import (
	"log"

	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
	"github.com/Lafetz/showdown-trivia-game/internal/repository"
	"github.com/Lafetz/showdown-trivia-game/internal/web"
	"github.com/Lafetz/showdown-trivia-game/internal/ws"
	"github.com/gorilla/sessions"
)

func main() {
	hashKey := "your-generated-hash-key"
	blockKey := "your-generated-block-key"
	store := sessions.NewCookieStore([]byte(hashKey), []byte(blockKey))
	repo := repository.NewStore()
	userservice := user.NewUserService(repo)
	hub := ws.NewHub()
	app := web.NewApp(userservice, hub, store)
	err := app.Run()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

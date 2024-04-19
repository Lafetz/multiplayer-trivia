package main

import (
	"flag"
	"log"

	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
	"github.com/Lafetz/showdown-trivia-game/internal/repository"
	"github.com/Lafetz/showdown-trivia-game/internal/web"
	"github.com/Lafetz/showdown-trivia-game/internal/web/ws"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	repo := &repository.Store{}
	userservice := user.NewUserService(repo)
	hub := ws.NewHub()
	app := web.NewApp(userservice, hub)
	err := app.Run()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

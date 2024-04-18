package main

import (
	"flag"
	"log"

	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
	"github.com/Lafetz/showdown-trivia-game/internal/repository"
	"github.com/Lafetz/showdown-trivia-game/internal/web"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	repo := &repository.Store{}
	userservice := user.NewUserService(repo)
	app := web.NewApp(userservice)
	err := app.Run()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

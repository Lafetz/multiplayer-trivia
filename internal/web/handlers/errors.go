package handlers

import (
	"log"
	"net/http"
	"runtime/debug"

	render "github.com/Lafetz/showdown-trivia-game/internal/web/Render"
)

func ServerError(w http.ResponseWriter, r *http.Request, err error, logger *log.Logger) {
	logger.Printf("%s\n%s", err.Error(), debug.Stack())
	err = render.InternalServer(w, r)
	if err != nil {
		log.Printf("%s\n%s", err.Error(), debug.Stack())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

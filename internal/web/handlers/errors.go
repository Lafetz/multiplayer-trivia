package handlers

import (
	"log"
	"net/http"
	"runtime/debug"

	render "github.com/Lafetz/showdown-trivia-game/internal/web/Render"
)

func ServerError(w http.ResponseWriter, r *http.Request, err error, logger *log.Logger) {
	w.WriteHeader(500)
	logger.Printf("%s\n%s", err.Error(), debug.Stack())
	err = render.DisplayToast(w, r, "Internal Server Error", true)
	if err != nil {
		log.Printf("%s\n%s", err.Error(), debug.Stack())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

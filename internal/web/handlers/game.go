package handlers

import (
	"errors"
	"log"
	"net/http"

	render "github.com/Lafetz/showdown-trivia-game/internal/web/Render"
	"github.com/Lafetz/showdown-trivia-game/internal/web/ws"
)

func Home(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, ok := r.Context().Value("username").(string)
		if !ok {
			ServerError(w, r, errors.New("username couldn't"), logger)
			return
		}

		err := render.Home(w, r, username)
		if err != nil {
			ServerError(w, r, err, logger)
		}

	}
}
func sendGamePage(w http.ResponseWriter, r *http.Request, create bool, id string) error {
	return render.SendGamePage(w, r, create, id)

}
func CreateGet(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gameOwner := true
		id := ""
		err := sendGamePage(w, r, gameOwner, id)
		if err != nil {
			ServerError(w, r, err, logger)
		}

	}
}
func ActiveGames(hub *ws.Hub, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rooms := hub.ListRooms()
		err := render.ActiveGames(w, r, rooms)
		if err != nil {
			ServerError(w, r, err, logger)
		}
	}
}
func Join(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gameOwner := false
		id := r.PathValue("id")
		err := sendGamePage(w, r, gameOwner, id)
		if err != nil {
			ServerError(w, r, err, logger)
		}
	}
}
func CreateWs(hub *ws.Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hub.CreateRoom(w, r)
	}
}
func JoinWs(hub *ws.Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		hub.JoinRoom(w, r, id)
	}
}

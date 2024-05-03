package handlers

import (
	"net/http"

	render "github.com/Lafetz/showdown-trivia-game/internal/web/Render"
	"github.com/Lafetz/showdown-trivia-game/internal/ws"
)

func Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, ok := r.Context().Value("username").(string)
		if !ok {
			render.InternalServer(w, r)
			return
		}

		render.Home(w, r, username)

	}
}
func sendGamePage(w http.ResponseWriter, r *http.Request, create bool, id string) {
	render.SendGamePage(w, r, create, id)

}
func CreateGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gameOwner := true
		id := ""
		sendGamePage(w, r, gameOwner, id)
	}
}
func ActiveGames(hub *ws.Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rooms := hub.ListRooms()
		render.ActiveGames(w, r, rooms)
	}
}
func Join() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gameOwner := false
		id := r.PathValue("id")
		sendGamePage(w, r, gameOwner, id)
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

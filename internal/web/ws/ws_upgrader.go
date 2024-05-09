package ws

import (
	"log"
	"net/http"

	"github.com/Lafetz/showdown-trivia-game/internal/core/entities"
	"github.com/google/uuid"
)

func (h *Hub) CreateRoom(w http.ResponseWriter, r *http.Request, username string, timer int, questions []entities.Question) {

	conn, err := WebsocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	id := uuid.New().String()[:7]

	room := NewRoom(id, username, timer, questions, h)
	h.addRoom(room)

	client := NewClient(conn, room, username)
	go client.readMessage()
	go client.writeMessage()
	room.addClient(client)

}
func (h *Hub) JoinRoom(w http.ResponseWriter, r *http.Request, roomId string, username string) {

	if roomId == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	room, err := h.getRoom(roomId)
	if err != nil {
		println(err)
		return
	}
	conn, err := WebsocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := NewClient(conn, room, username)
	go client.readMessage()
	go client.writeMessage()
	room.addClient(client)

}

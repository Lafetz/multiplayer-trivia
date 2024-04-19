package ws

import (
	"log"
	"net/http"
)

func (h *Hub) CreateRoom(w http.ResponseWriter, r *http.Request) {

	conn, err := WebsocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	room := NewRoom("test")
	h.addRoom(room)

	if err != nil {
		println(err)

		return
	}

	client := NewClient(conn, room)
	go client.readMessage()
	go client.writeMessage()

}
func (h *Hub) JoinRoom(w http.ResponseWriter, r *http.Request) {
	roomId := r.URL.Query().Get("id")
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
	client := NewClient(conn, room)
	go client.readMessage()
	go client.writeMessage()
}

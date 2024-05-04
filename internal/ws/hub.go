package ws

import (
	"errors"
	"net/http"
	"sync"

	"github.com/Lafetz/showdown-trivia-game/internal/web/entity"
	"github.com/gorilla/websocket"
)

var (
	WebsocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     checkOrigin,
	}
)

type Hub struct {
	rooms RoomList
	sync.RWMutex
}

func (h *Hub) getRoom(roomId string) (*Room, error) {
	if room, ok := h.rooms[roomId]; ok {
		return room, nil
	} else {
		return nil, errors.New(("room doens't exist"))
	}
}
func (h *Hub) addRoom(room *Room) {
	h.Lock()
	defer h.Unlock()
	h.rooms[room.Id] = room
}
func (h *Hub) ListRooms() []entity.RoomData {
	h.Lock()
	defer h.Unlock()
	var rooms []entity.RoomData
	for _, r := range h.rooms {

		rooms = append(rooms, entity.RoomData{
			Owner:   r.owner,
			Id:      r.Id,
			Players: r.getUsers(),
		})
	}
	return rooms
}
func (h *Hub) removeRoom(room *Room) {
	h.Lock()
	defer h.Unlock()
	delete(h.rooms, room.Id)
}

func NewHub() *Hub {
	h := &Hub{
		rooms: make(RoomList),
		// handlers: make(map[string]EventHandler),
	}
	return h
}
func checkOrigin(r *http.Request) bool {
	// origin := r.Header.Get("Origin")
	// switch origin {
	// case "http://localhost:300":
	// 	return true
	// default:
	// 	return false
	// }
	//	id := uuid.New().String()[:7]

	return true
}

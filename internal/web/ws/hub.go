package ws

import (
	"errors"
	"log"
	"log/slog"
	"net/http"
	"sync"

	webentities "github.com/Lafetz/showdown-trivia-game/internal/web/entity"
	"github.com/gorilla/websocket"
)

var (
	WebsocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     checkOrigin,
	}
)
var (
	ErrRoomNotExist = errors.New("room doens't exist")
)

type Hub struct {
	Logger *slog.Logger
	rooms  RoomList
	sync.RWMutex
}

func (h *Hub) getRoom(roomId string) (*Room, error) {
	if room, ok := h.rooms[roomId]; ok {
		return room, nil
	} else {
		return nil, ErrRoomNotExist
	}
}
func (h *Hub) addRoom(room *Room) {
	h.Lock()
	defer h.Unlock()
	h.rooms[room.Id] = room
}
func (h *Hub) ListRooms() []webentities.RoomData {
	h.Lock()
	defer h.Unlock()
	var rooms []webentities.RoomData
	for _, r := range h.rooms {
		if r.Game.GameStarted {
			continue
		}
		rooms = append(rooms, webentities.RoomData{
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
	if _, ok := h.rooms[room.Id]; !ok {
		log.Printf("attempted to remove non-existent room: %s", room.Id)
		return
	}
	delete(h.rooms, room.Id)
}

func NewHub(logger *slog.Logger) *Hub {
	h := &Hub{
		rooms:  make(RoomList),
		Logger: logger,
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

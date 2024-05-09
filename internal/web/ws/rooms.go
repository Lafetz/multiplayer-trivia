package ws

import (
	"sync"
	"time"

	"github.com/Lafetz/showdown-trivia-game/internal/core/entities"
	"github.com/Lafetz/showdown-trivia-game/internal/core/game"
	render "github.com/Lafetz/showdown-trivia-game/internal/web/Render"
)

type RoomList map[string]*Room
type Room struct {
	hub     *Hub
	clients ClientList
	Game    game.Game
	Id      string
	owner   string
	sync.RWMutex
}

func (r *Room) sendMsg(msg []byte) {

	for c := range r.clients {

		c.egress <- msg
	}

}

func NewRoom(id string, owner string, timer int, questions []entities.Question, hub *Hub) *Room {
	time := time.Duration(timer) * time.Second

	r := &Room{
		Id:      id,
		clients: make(ClientList),
		owner:   owner,
		Game:    *game.NewGame(questions, time),
		hub:     hub,
	}
	go func() {
		for m := range r.Game.Message {
			switch m.MsgType {
			case game.MsgQuestion:
				if payload, ok := m.Payload.(entities.Question); ok {
					buff := render.RenderQuestion(payload, r.Game.CurrentQues, len(r.Game.Questions), timer, r.Game.Players)
					r.sendMsg(buff.Bytes())
				} else {
					continue
				}
			case game.MsgInfo:
				if payload, ok := m.Payload.(game.Info); ok {
					buff := render.RenderGameMessage(payload)
					r.sendMsg(buff.Bytes())
				} else {
					continue
				}
			case game.MsgGameEnd:
				if payload, ok := m.Payload.(game.Winners); ok {
					buff := render.GameEnd(payload)
					r.sendMsg(buff.Bytes())
					r.hub.removeRoom(r)
				} else {
					continue
				}
			}

		}
	}()
	return r
}
func (r *Room) addClient(client *Client) {
	r.Lock()
	defer r.Unlock()
	r.clients[client] = true
	buff := render.RenderPlayers(r.Id, r.getUsers())
	r.sendMsg(buff.Bytes())
	//

	//
}
func (r *Room) removeClient(client *Client) {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.clients[client]; ok {
		client.connection.Close()
		delete(r.clients, client)
	}
	if len(r.clients) == 0 {
		if _, ok := r.hub.rooms[r.Id]; ok {
			r.hub.removeRoom(r)
			return
		}
	}
}
func (r *Room) getUsers() []string {
	var users []string
	for c := range r.clients {
		users = append(users, c.Username)
	}
	return users
}

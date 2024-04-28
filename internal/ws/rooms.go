package ws

import (
	"sync"

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

func NewRoom(id string) *Room {
	questions := []game.Question{
		{Question: "What is 2+2?", Options: []string{"A. 2", "B. 4", "C. 43", "D. 1"}, CorrectAnswer: "B"},
		{Question: "What is the capital of France?", Options: []string{"A. London", "B. Berlin", "C. Paris", "D. Rome"}, CorrectAnswer: "C"},
	}
	g := *game.NewGame(questions)
	r := &Room{
		Id:      id,
		clients: make(ClientList),
		owner:   "unkownd_owner",
		Game:    g,
	}
	go func() {
		for m := range g.Message {
			switch m.MsgType {
			case "question":
				if payload, ok := m.Payload.(game.Question); ok {
					buff := render.RenderQuestion(payload, "")
					r.sendMsg(buff.Bytes())
				} else {
					continue
				}
			case "info":
				if payload, ok := m.Payload.(game.Info); ok {
					buff := render.RenderGameMessage(payload)
					r.sendMsg(buff.Bytes())
				} else {
					continue
				}
			case "game_end":
				if payload, ok := m.Payload.(game.Winners); ok {
					buff := render.GameEnd(payload)
					r.sendMsg(buff.Bytes())
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
}
func (r *Room) removeClient(client *Client) {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.clients[client]; ok {
		client.connection.Close()
		delete(r.clients, client)
	}
}
func (r *Room) getUsers() []string {
	var users []string
	for c := range r.clients {
		users = append(users, c.Username)
	}
	return users
}

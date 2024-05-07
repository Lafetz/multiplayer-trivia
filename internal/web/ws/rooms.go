package ws

import (
	"sync"

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

func NewRoom(id string) *Room {
	questions := []entities.Question{
		{Question: "What is 2+2?", Options: []string{"3", "4", "5", "6"}, CorrectAnswer: "4"},
		{Question: "What is the capital of France?", Options: []string{"London", "Berlin", "Paris", "Rome"}, CorrectAnswer: "Paris"},
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
			case game.MsgQuestion:
				if payload, ok := m.Payload.(entities.Question); ok {
					buff := render.RenderQuestion(payload, r.Game.CurrentQues, len(g.Questions))
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

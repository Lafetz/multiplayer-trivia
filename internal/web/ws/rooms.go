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
					buff, err := render.RenderQuestion(payload, r.Game.CurrentQues, len(r.Game.Questions), timer, r.Game.Players)
					if err != nil {
						r.hub.Logger.Error(err.Error())
					}
					r.sendMsg(buff.Bytes())
				} else {
					continue
				}
			case game.MsgInfo:
				if payload, ok := m.Payload.(game.Info); ok {
					buff, err := render.RenderGameMessage(payload)
					if err != nil {
						r.hub.Logger.Error(err.Error())
					}
					r.sendMsg(buff.Bytes())
				} else {
					continue
				}
			case game.MsgGameEnd:
				if payload, ok := m.Payload.(game.Winners); ok {
					buff, err := render.GameEnd(payload)
					if err != nil {
						r.hub.Logger.Error(err.Error())
					}
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
	r.hub.m.WebsocketConns.Inc()
	buff, err := render.RenderPlayers(r.Id, r.getUsers())
	if err != nil {
		r.hub.Logger.Error(err.Error())
	}
	r.sendMsg(buff.Bytes())
}
func (r *Room) removeClient(client *Client) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.clients[client]; ok {
		err := client.connection.Close()
		if err != nil {
			return err
		}
		delete(r.clients, client)
		r.hub.m.WebsocketConns.Dec()
		//
	}
	if len(r.clients) == 0 { // remove room if there a no connected users
		if _, ok := r.hub.rooms[r.Id]; ok {
			r.hub.removeRoom(r)
			return nil
		}
	}
	return nil
}
func (r *Room) getUsers() []string {
	var users []string
	for c := range r.clients {
		users = append(users, c.Username)
	}
	return users
}

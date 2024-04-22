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
	println(len(r.clients))
	for c := range r.clients {
		print(c.Username)
		c.egress <- msg
	}
	println("done sending pe")
}

//	func (r *Room) sendMsg(msg string) {
//		println(len(r.clients))
//		for c := range r.clients {
//			print(c.Username)
//			c.egress <- []byte(msg)
//		}
//		println("done sending pe")
//	}
func NewRoom(id string) *Room {
	questions := []game.Question{
		{Question: "What is 2+2?", Options: []string{"3", "4", "5", "6"}, CorrectAnswer: "B"},
		{Question: "What is the capital of France?", Options: []string{"A) London", "B) Berlin", "C) Paris", "D) Rome"}, CorrectAnswer: "C"},
	}
	g := *game.NewGame(questions)
	r := &Room{
		Id:      id,
		clients: make(ClientList),
		owner:   "unkownd_owner",
		Game:    g,
	}
	go func() {
		for q := range g.Message {
			buff := render.RenderQuestion(q)
			println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
			r.sendMsg(buff.Bytes())
		}
	}()
	return r
}
func (r *Room) addClient(client *Client) {
	r.Lock()
	defer r.Unlock()
	r.clients[client] = true
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

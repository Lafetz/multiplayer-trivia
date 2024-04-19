package ws

import "sync"

type RoomList map[string]*Room
type Room struct {
	hub     *Hub
	clients ClientList
	Id      string
	sync.RWMutex
}

func NewRoom(name string) *Room {
	return &Room{
		Id:      name,
		clients: make(ClientList),
	}
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

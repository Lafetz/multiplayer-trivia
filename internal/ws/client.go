package ws

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var (
	pongWait     = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type ClientList map[*Client]bool
type Client struct {
	connection *websocket.Conn
	room       *Room
	egress     chan []byte
}

func NewClient(conn *websocket.Conn, room *Room) *Client {
	c := &Client{
		connection: conn,
		room:       room,
		egress:     make(chan []byte),
	}
	room.addClient(c)
	return c
}
func (c *Client) readMessage() {
	defer func() {
		c.room.removeClient(c)
	}()
	if err := c.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println(err)
		return
	}
	c.connection.SetReadLimit(512)
	c.connection.SetPongHandler(c.pongHandler)
	for {
		_, payload, err := c.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure) {
				log.Println("error reading msg: ", err)
			}
			break
		}
		var req Event
		if err := json.Unmarshal(payload, &req); err != nil {
			log.Println("error reading message", err)
			break
		}
		println(req.EventType, req.Payload)
		for x := range c.room.clients {
			x.egress <- payload
		}
		// if err := c.hub.routeMessages(req, c); err != nil {
		// 	log.Println("error routing message", err)
		// 	break
		// }
	}
}
func (c *Client) writeMessage() {
	defer func() {
		c.room.removeClient(c)
	}()
	ticker := time.NewTicker(pingInterval)
	for {
		select {
		case message, ok := <-c.egress:

			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection is closed", err)
				}
				return
			}

			data, err := json.Marshal(message)
			if err != nil {
				log.Print(err)
				return
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println("failed to send msg", err)
			}
			log.Print("message sent")
		case <-ticker.C:
			log.Println("ping")
			if err := c.connection.WriteMessage(websocket.PingMessage, []byte("")); err != nil {
				log.Println("ping failed", err)
				return
			}
		}
	}
}
func (c *Client) pongHandler(pongMsg string) error {
	log.Println("pong")
	return c.connection.SetReadDeadline(time.Now().Add(pongWait))
}

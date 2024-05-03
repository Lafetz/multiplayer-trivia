package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Lafetz/showdown-trivia-game/internal/core/game"
	render "github.com/Lafetz/showdown-trivia-game/internal/web/Render"
	"github.com/gorilla/websocket"
)

var (
	pongWait     = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type ClientList map[*Client]bool
type Client struct {
	Username   string
	connection *websocket.Conn
	room       *Room
	egress     chan []byte
}

func NewClient(conn *websocket.Conn, room *Room) *Client {
	c := &Client{
		Username:   fmt.Sprintf("user%s", strconv.Itoa(len(room.clients))),
		connection: conn,
		room:       room,
		egress:     make(chan []byte),
	}

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
		_, msg, err := c.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure) {
				log.Println("error reading msg: ", err)
			}
			break
		}

		var req Event
		if err := json.Unmarshal(msg, &req); err != nil {
			log.Println("error reading message", err)
			break
		}
		//check if valid
		//
		switch req.EventType {
		case StartGame:
			var players []*game.Player
			for c := range c.room.clients {
				println(c.Username, "hmmm")
				players = append(players, game.NewPlayer(c.Username))
			}
			go c.room.Game.Start(players)
		case SendAnswer:
			answer := game.NewAnswer(c.Username, req.Payload)
			c.room.Game.AnswerCh <- answer
			buff := render.RenderUserAnswer(req.Payload)
			c.egress <- buff.Bytes()
		default:
			log.Println("unknown event type:", req.EventType)
		}

	}
}
func (c *Client) writeMessage() {
	defer func() {
		c.room.removeClient(c)
	}()
	ticker := time.NewTicker(pingInterval)
	//
	// buf := render.RenderWS()
	// err := c.connection.WriteMessage(websocket.TextMessage, buf.Bytes())
	// println("ping")
	// if err != nil {
	// 	return
	// }
	//
	for {
		select {
		case message, ok := <-c.egress:

			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection is closed", err)
				}
				return
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println("failed to send msg", err)
			}
			log.Print("message sent")

		case <-ticker.C:

			if err := c.connection.WriteMessage(websocket.PingMessage, []byte("")); err != nil {
				log.Println("ping failed", err)
				return
			}
		}
	}
}
func (c *Client) pongHandler(pongMsg string) error {

	return c.connection.SetReadDeadline(time.Now().Add(pongWait))
}

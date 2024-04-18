package ws

import (
	"encoding/json"
	"fmt"
)

type Event struct {
	EventType string          `json:"type"`
	Payload   json.RawMessage `json:"payload"`
}
type EventHandler func(event Event, c *Client) error

const (
	EventSendMessage = "send_message"
)

type SendMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
}

func sendMessage(event Event, c *Client) error {
	fmt.Println(event.EventType, event.Payload)
	return nil
}

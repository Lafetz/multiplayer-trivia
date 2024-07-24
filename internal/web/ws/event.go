package ws

type Event struct {
	EventType string `json:"type"`
	Payload   string `json:"payload"`
}

// type EventHandler func(event Event, c *Client) error

const (
	EventSendMessage = "send_message"
	NewPlayer        = "new_player"
	StartGame        = "start_game"
	SendAnswer       = "send_answer"
)

// type SendMessageEvent struct {
// 	Message string `json:"message"`
// 	From    string `json:"from"`
// }

// func sendMessage(event Event, c *Client) error {
// 	fmt.Println(event.EventType, event.Payload)
// 	return nil
// }

package game

const (
	MsgGameEnd  = "game_end"
	MsgQuestion = "question"
	MsgInfo     = "info"
)

type Message struct {
	MsgType string
	Payload interface{}
}

func NewMessage(MsgType string, payload interface{}) Message {
	return Message{
		MsgType: MsgType,
		Payload: payload,
	}
}

type Player struct {
	Username string
	Score    int
}

func NewPlayer(username string) *Player {
	return &Player{
		Username: username,
		Score:    0,
	}
}

type Answer struct {
	username string
	answer   string
}

func NewAnswer(username string, ans string) Answer {
	return Answer{
		username: username,
		answer:   ans,
	}
}

type Info struct {
	InfoType string
	Text     string
}
type Winners map[string]int

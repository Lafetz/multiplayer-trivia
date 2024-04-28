package game

import (
	"time"
)

const (
	DefaultTimerSpan = 5 * time.Second
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

type Game struct {
	Questions   []Question
	Players     []*Player
	CurrentQues int
	AnswerCh    chan Answer
	Message     chan Message
	timerSpan   time.Duration
	gameStarted bool
}

func (g *Game) Start(players []*Player) {
	g.Players = players
	g.gameStarted = true
	for _, question := range g.Questions {
		g.AskQuestion(question)
		time.Sleep(g.timerSpan)
	}
	g.DisplayWinner()
}

func (g *Game) AskQuestion(question Question) {
	timer := time.NewTimer(g.timerSpan)
	g.Message <- NewMessage("question", question)
	go func() {
		for {
			select {
			case answer := <-g.AnswerCh:
				// println("some one sent answer whooooooooooooo")
				if answer.answer == question.CorrectAnswer {
					for _, player := range g.Players {
						if player.Username == answer.username {
							player.Score++
						}
					}

				}
			case <-timer.C:
				g.CurrentQues++

				return
			}
		}
	}()
}
func (g *Game) DisplayWinner() {

	winners := make(Winners)
	for _, player := range g.Players {
		winners[player.Username] = player.Score
	}
	g.Message <- NewMessage("game_end", winners)
}
func NewGame(questions []Question) *Game {
	return &Game{
		Questions: questions,
		//	Players:   players,
		AnswerCh:    make(chan Answer),
		Message:     make(chan Message),
		timerSpan:   DefaultTimerSpan,
		CurrentQues: 0,
		gameStarted: false,
	}
}

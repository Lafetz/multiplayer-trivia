package game

import (
	"sync"
	"time"
)

const (
	DefaultTimerSpan = 1 * time.Second
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
	//	playerAnswers map[int][]Answer
	sync.RWMutex
}

func (g *Game) Start(players []*Player) {

	g.Players = players
	g.gameStarted = true
	for _, question := range g.Questions {
		doneCh := make(chan struct{})
		g.CurrentQues++
		g.AskQuestion(question, doneCh)

	}

}

func (g *Game) AskQuestion(question Question, doneCh chan struct{}) {

	timer := time.NewTimer(g.timerSpan)

	g.Message <- NewMessage("question", question)

	answers := make(map[string]string)

	for {

		select {
		case answer := <-g.AnswerCh:
			answers[answer.username] = answer.answer

		case <-timer.C:
			for username, answer := range answers {

				if question.CorrectAnswer == answer {

					for _, player := range g.Players {

						if player.Username == username {

							player.Score++
						}
					}

				}

			}

			if len(g.Questions) == (g.CurrentQues) {
				g.DisplayWinner()
			}
			answers = make(map[string]string)
			close(doneCh)
		case <-doneCh:
			return
		default:
			continue

		}
	}
	//}()
}
func (g *Game) DisplayWinner() {

	winners := make(Winners)
	for _, player := range g.Players {
		winners[player.Username] = player.Score
	}
	g.Message <- NewMessage("game_end", winners)
	println("closed")
	close(g.Message)
}
func NewGame(questions []Question) *Game {
	return &Game{
		Questions:   questions,
		AnswerCh:    make(chan Answer),
		Message:     make(chan Message),
		timerSpan:   DefaultTimerSpan,
		CurrentQues: 0,
		gameStarted: false,
	}
}

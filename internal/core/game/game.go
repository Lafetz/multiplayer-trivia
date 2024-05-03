package game

import (
	"time"
)

const (
	DefaultTimerSpan = 3 * time.Second
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
	Questions     []Question
	Players       []*Player
	CurrentQues   int
	AnswerCh      chan Answer
	Message       chan Message
	timerSpan     time.Duration
	gameStarted   bool
	playerAnswers map[int][]Answer
}

func (g *Game) Start(players []*Player) {
	g.Players = players
	g.gameStarted = true
	for _, question := range g.Questions {
		g.CurrentQues++
		g.AskQuestion(question)
		//time.Sleep(g.timerSpan)

	}

}

func (g *Game) AskQuestion(question Question) {
	timer := time.NewTimer(g.timerSpan)
	g.Message <- NewMessage("question", question)
	println(g.CurrentQues)
	//go func() {

	answers := make(map[string]string)
	for {
		select {
		case answer := <-g.AnswerCh:
			answers[answer.username] = answer.answer
		case <-timer.C:

			for username, answer := range answers {
				println(username, answer)
				if question.CorrectAnswer == answer {
					for _, player := range g.Players {
						if player.Username == username {

							player.Score++
						}
					}

				}

				if len(g.Questions) == (g.CurrentQues) {
					g.DisplayWinner()
				}

				return
			}

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
}
func NewGame(questions []Question) *Game {
	return &Game{
		Questions: questions,
		//	Players:   players,
		AnswerCh:      make(chan Answer),
		Message:       make(chan Message),
		timerSpan:     DefaultTimerSpan,
		CurrentQues:   0,
		gameStarted:   false,
		playerAnswers: make(map[int][]Answer),
	}
}

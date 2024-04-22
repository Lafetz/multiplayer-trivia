package game

import (
	"fmt"
	"time"
)

const (
	DefaultTimerSpan = 10 * time.Second
)

type Message struct {
	MsgType string
	payload interface{}
}

type Game struct {
	Questions   []Question
	Players     []*Player
	CurrentQues int
	AnswerCh    chan Answer
	Message     chan Question
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
	g.Message <- question
	go func() {
		for {
			select {
			case answer := <-g.AnswerCh:
				if answer.answer == question.CorrectAnswer {
					for _, player := range g.Players {
						if player.Username == answer.username {
							player.Score++
						}
					}

				}
			case <-timer.C:
				g.CurrentQues++
				fmt.Println("Time's up!")
				return
			}
		}
	}()
}
func (g *Game) DisplayWinner() {
	fmt.Println("Game Over! Final Scores:")
	for _, player := range g.Players {
		fmt.Printf("Player %s: %d\n", player.Username, player.Score)
	}
}
func NewGame(questions []Question) *Game {
	return &Game{
		Questions: questions,
		//	Players:   players,
		AnswerCh:    make(chan Answer),
		Message:     make(chan Question),
		timerSpan:   DefaultTimerSpan,
		gameStarted: false,
	}
}

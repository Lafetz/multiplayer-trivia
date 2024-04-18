package game

import (
	"fmt"
	"time"
)

const (
	DefaultTimerSpan = 10 * time.Second
)

type Game struct {
	Questions   []Question
	Players     []*Player
	CurrentQues int
	AnswerCh    chan Answer
	message     chan string
	timerSpan   time.Duration
}

func (g *Game) Start() {
	for _, question := range g.Questions {
		g.AskQuestion(question)
		time.Sleep(g.timerSpan)
	}
	g.DisplayWinner()
}

func (g *Game) AskQuestion(question Question) {
	timer := time.NewTimer(g.timerSpan)
	g.message <- fmt.Sprintf("questions %s ", question)
	go func() {
		for {
			select {
			case answer := <-g.AnswerCh:
				if answer.answer == question.CorrectAnswer {
					for _, player := range g.Players {
						if player.username == answer.username {
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
		fmt.Printf("Player %s: %d\n", player.username, player.Score)
	}
}
func newGame(questions []Question, players []*Player) *Game {
	return &Game{
		Questions: questions,
		Players:   players,
		AnswerCh:  make(chan Answer),
		message:   make(chan string),
		timerSpan: DefaultTimerSpan,
	}
}

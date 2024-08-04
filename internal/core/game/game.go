package game

import (
	"sync"
	"time"

	"github.com/Lafetz/showdown-trivia-game/internal/core/entities"
)

const (
	DefaultTimerSpan = 3 * time.Second
)

type Game struct {
	Questions   []entities.Question
	Players     []*Player
	CurrentQues int
	AnswerCh    chan Answer
	Message     chan Message
	timerSpan   time.Duration
	GameStarted bool
	sync.RWMutex
}

func (g *Game) Start(players []*Player) {

	g.Players = players
	g.GameStarted = true
	for _, question := range g.Questions {
		doneCh := make(chan struct{})
		g.CurrentQues++
		g.AskQuestion(question, doneCh)

	}

}
func (g *Game) endOfQuestion(question entities.Question, answers map[string]string) { //score correct answers
	for username, answer := range answers {
		if question.CorrectAnswer == answer {
			g.score(username)
		}

	}
}

func (g *Game) score(username string) {
	for _, player := range g.Players {
		if player.Username == username {

			player.Score++

		}
	}
}
func (g *Game) AskQuestion(question entities.Question, doneCh chan struct{}) {

	timer := time.NewTimer(g.timerSpan)

	g.Message <- NewMessage(MsgQuestion, question)

	answers := make(map[string]string)

	for {

		select {
		case answer := <-g.AnswerCh:
			answers[answer.username] = answer.answer

		case <-timer.C:
			g.endOfQuestion(question, answers)

			if len(g.Questions) == (g.CurrentQues) { //was the last question
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

}
func (g *Game) DisplayWinner() {

	winners := make(Winners)
	for _, player := range g.Players {
		winners[player.Username] = player.Score
	}
	g.Message <- NewMessage(MsgGameEnd, winners)

	close(g.Message)
}
func NewGame(questions []entities.Question, timer time.Duration) *Game {

	return &Game{
		Questions:   questions,
		AnswerCh:    make(chan Answer),
		Message:     make(chan Message),
		timerSpan:   timer,
		CurrentQues: 0,
		GameStarted: false,
	}
}

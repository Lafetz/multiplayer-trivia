package game

import (
	"testing"
	"time"

	"github.com/Lafetz/showdown-trivia-game/internal/core/entities"
)

func TestGame(t *testing.T) {
	questions := []entities.Question{
		{Question: "What is 2+2?", Options: []string{"A) 3", "B) 4", "C) 5", "D) 6"}, CorrectAnswer: "B"},
		{Question: "What is the capital of France?", Options: []string{"A) London", "B) Berlin", "C) Paris", "D) Rome"}, CorrectAnswer: "C"},
	}
	players := []*Player{
		NewPlayer("panzer"),
		NewPlayer("leopard"),
	}
	game := NewGame(questions)
	//change default timer so that tests finishes quicker
	span := 100 * time.Microsecond
	game.timerSpan = span

	// start the game
	go game.Start(players)

loop:
	for m := range game.Message {
		switch m.MsgType {
		case "question":
			if game.CurrentQues == 1 {
				game.AnswerCh <- NewAnswer("panzer", "B")
				game.AnswerCh <- NewAnswer("leopard", "B")
			} else {

				game.AnswerCh <- NewAnswer("panzer", "C")
				game.AnswerCh <- NewAnswer("leopard", "B")
				//continue
			}
		case "info":
			continue
		case "game_end":
			break loop
		}

	}

	// verify player scores
	expectedScores := []int{2, 1} // expected scores for each player
	for i, player := range game.Players {

		if player.Score != expectedScores[i] {
			t.Errorf("Player %s: got score %d, expected %d", player.Username, player.Score, expectedScores[i])
		}
	}
}

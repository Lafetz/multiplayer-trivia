package game

import (
	"testing"
	"time"

	"github.com/Lafetz/showdown-trivia-game/internal/core/entities"
)

func TestGame(t *testing.T) {
	questions := []entities.Question{
		{Question: "What is 2+2?", Options: []string{"3", "4", "5", "6"}, CorrectAnswer: "4"},
		{Question: "What is the capital of France?", Options: []string{"London", "Berlin", "Paris", "Rome"}, CorrectAnswer: "Paris"},
	}
	players := []*Player{
		NewPlayer("panzer"),
		NewPlayer("leopard"),
	}
	span := 100 * time.Microsecond
	game := NewGame(questions, span)
	//change default timer so that tests finishes quicker
	//span := 100 * time.Microsecond
	//game.timerSpan = span

	// start the game
	go game.Start(players)

loop:
	for m := range game.Message {
		switch m.MsgType {
		case "question":
			if game.CurrentQues == 1 {
				game.AnswerCh <- NewAnswer("panzer", "4")
				game.AnswerCh <- NewAnswer("leopard", "4")
			} else {

				game.AnswerCh <- NewAnswer("panzer", "Paris")
				game.AnswerCh <- NewAnswer("leopard", "London")
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

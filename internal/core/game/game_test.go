package game

import (
	"testing"
	"time"
)

func TestGame(t *testing.T) {
	questions := []Question{
		{Question: "What is 2+2?", Options: []string{"A) 3", "B) 4", "C) 5", "D) 6"}, CorrectAnswer: "B"},
		{Question: "What is the capital of France?", Options: []string{"A) London", "B) Berlin", "C) Paris", "D) Rome"}, CorrectAnswer: "C"},
	}
	players := []*Player{
		NewPlayer("panzer"),
		NewPlayer("leopard"),
	}
	game := NewGame(questions)
	//change default timer so that tests finishes quicker
	span := 500 * time.Microsecond
	game.timerSpan = span
	// go func() {
	// 	for msg := range game.Message {
	// 		//println(msg)
	// 	}
	// }()
	// start the game
	go game.Start(players)
	// simulate players answering questions
	go func() {
		game.AnswerCh <- NewAnswer("panzer", "1")
		game.AnswerCh <- NewAnswer("leopard", "1")
		time.Sleep(span)
		game.AnswerCh <- NewAnswer("panzer", "2") //playe a answers first
		time.Sleep(span / 100)
		game.AnswerCh <- NewAnswer("leopard", "2")
	}()

	// wait for the game to finish
	<-time.After(span * 1000)

	// verify player scores
	expectedScores := []int{1, 2} // expected scores for each player
	for i, player := range game.Players {
		if player.Score != expectedScores[i] {
			t.Errorf("Player %s: got score %d, expected %d", player.Username, player.Score, expectedScores[i])
		}
	}
}

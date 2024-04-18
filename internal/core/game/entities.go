package game

type Question struct {
	Question      string
	Options       []string
	CorrectAnswer string
}

func NewQuestion() {

}

type Player struct {
	username string
	Score    int
}

func NewPlayer(username string) *Player {
	return &Player{
		username: username,
	}
}

type Answer struct {
	username string
	answer   string
}

func newAnswer(username string, ans string) Answer {
	return Answer{
		username: username,
		answer:   ans,
	}
}

package game

type Question struct {
	Question      string
	Options       []string
	CorrectAnswer string
}

func NewQuestion() {

}

type Player struct {
	Username string
	Score    int
}

func NewPlayer(username string) *Player {
	return &Player{
		Username: username,
		Score:    0,
	}
}

type Answer struct {
	username string
	answer   string
}

func NewAnswer(username string, ans string) Answer {
	return Answer{
		username: username,
		answer:   ans,
	}
}

type Info struct {
	InfoType string
	Text     string
}
type Winners map[string]int

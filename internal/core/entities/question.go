package entities

type Question struct {
	Question      string
	Options       []string
	CorrectAnswer string
}

func NewQuestion(question string, options []string, correctAnswer string) Question {
	return Question{
		Question:      question,
		Options:       options,
		CorrectAnswer: correctAnswer,
	}
}

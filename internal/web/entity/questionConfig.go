package webentities

type QuestionConfig struct {
	Category int
	Timer    int
	Amount   int
}

func NewQuestionConfig(category int, timer int, amount int) QuestionConfig {
	return QuestionConfig{
		Category: category,
		Timer:    timer,
		Amount:   amount,
	}
}

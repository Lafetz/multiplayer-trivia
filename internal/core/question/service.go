package question

import "github.com/Lafetz/showdown-trivia-game/internal/core/entities"

type QuestionService struct {
	questionClient QuestionClientApi
}

func (q *QuestionService) GetQuestions(amount int, category int) ([]entities.Question, error) {
	return q.questionClient.GetQuestions(amount, category)
}
func (q *QuestionService) GetCategories() ([]Category, error) {
	return q.questionClient.GetCategories()
}

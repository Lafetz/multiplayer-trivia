package question

import "github.com/Lafetz/showdown-trivia-game/internal/core/entities"

type QuestionServiceApi interface {
	GetQuestions(amount int, category int) ([]entities.Question, error)
	GetCategories() ([]Category, error)
}
type QuestionClientApi interface {
	GetQuestions(amount int, category int) ([]entities.Question, error)
	GetCategories() ([]Category, error)
}

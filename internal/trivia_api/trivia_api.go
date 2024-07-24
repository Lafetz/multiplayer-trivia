package triviaapi

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/Lafetz/showdown-trivia-game/internal/core/entities"
	"github.com/Lafetz/showdown-trivia-game/internal/core/question"
)

type triviaClient struct {
	baseURL string
}

func (t *triviaClient) GetQuestions(amount int, category int) ([]entities.Question, error) {
	apiURL := fmt.Sprintf("%s/api.php?amount=%d&type=multiple&category=%d", t.baseURL, amount, category)
	questions, err := fetchQuestions(apiURL)
	if err != nil {
		return nil, err
	}
	var questionsDomain []entities.Question

	for _, tq := range questions {
		question := ConvertToQuestion(tq)
		questionsDomain = append(questionsDomain, question)
	}
	return questionsDomain, nil
}

func (t *triviaClient) GetCategories() ([]question.Category, error) {
	url := fmt.Sprintf("%s/api_category.php", t.baseURL) //https://opentdb.com/
	return fetchCategories(url)
}
func NewTriviaClient() *triviaClient {
	return &triviaClient{
		baseURL: "https://opentdb.com",
	}
}

// helper functions
type categoryResponse struct {
	TriviaCategories []question.Category `json:"trivia_categories"`
}

func fetchCategories(url string) ([]question.Category, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// Check response status code
	if response.StatusCode != http.StatusOK {
		return nil, err
	}
	defer response.Body.Close()
	// Decode JSON response
	var categoryResponse categoryResponse
	err = json.NewDecoder(response.Body).Decode(&categoryResponse)
	if err != nil {
		return nil, err
	}

	return categoryResponse.TriviaCategories, nil
}
func fetchQuestions(apiURL string) ([]triviaQuestion, error) {
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, err
	}
	var jsonResponse struct {
		ResponseCode int              `json:"response_code"`
		Results      []triviaQuestion `json:"results"`
	}
	if err := json.NewDecoder(response.Body).Decode(&jsonResponse); err != nil {
		return nil, err
	}

	if jsonResponse.ResponseCode != 0 {
		return nil, NewAPIError(jsonResponse.ResponseCode)
	}
	return jsonResponse.Results, nil
}
func shuffleOptions(options []string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := len(options)
	for i := n - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		options[i], options[j] = options[j], options[i]
	}

}
func ConvertToQuestion(triviaQuestionData triviaQuestion) entities.Question {
	question := triviaQuestionData.Question
	incorrectAnswers := triviaQuestionData.IncorrectAnswers
	correctAnswer := triviaQuestionData.CorrectAnswer
	options := append(incorrectAnswers, correctAnswer)
	shuffleOptions(options)
	return entities.NewQuestion(question, options, correctAnswer)
}

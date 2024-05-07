package triviaapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Lafetz/showdown-trivia-game/internal/core/entities"
	"github.com/Lafetz/showdown-trivia-game/internal/core/question"
)

func TestConvertToQuestion(t *testing.T) {

	testCases := []struct {
		Name           string
		Input          triviaQuestion
		ExpectedOutput entities.Question
	}{
		{
			Name: "Basic Conversion",
			Input: triviaQuestion{
				Type:             "multiple",
				Difficulty:       "easy",
				Category:         "Science: Computers",
				Question:         "What does CPU stand for?",
				CorrectAnswer:    "Central Processing Unit",
				IncorrectAnswers: []string{"Computer Personal Unit", "Central Personal Unit", "Central Processor Unit"},
			},
			ExpectedOutput: entities.Question{
				Question:      "What does CPU stand for?",
				CorrectAnswer: "Central Processing Unit",
				Options:       []string{"Computer Personal Unit", "Central Personal Unit", "Central Processor Unit", "Central Processing Unit"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			actualOutput := ConvertToQuestion(tc.Input)

			if actualOutput.Question != tc.ExpectedOutput.Question {
				t.Errorf("Test case '%s' failed: unexpected question", tc.Name)
			}
			if actualOutput.CorrectAnswer != tc.ExpectedOutput.CorrectAnswer {
				t.Errorf("Test case '%s' failed: unexpected correct answer", tc.Name)
			}

			expectedOptionsSet := make(map[string]bool)
			for _, opt := range tc.ExpectedOutput.Options {
				expectedOptionsSet[opt] = true
			}
			for _, opt := range actualOutput.Options {
				if !expectedOptionsSet[opt] {
					t.Errorf("Test case '%s' failed: unexpected option '%s'", tc.Name, opt)
				}
			}

		})
	}
}

func TestFetchQuestions(t *testing.T) {
	// Mock server to simulate API responses
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Determine which API endpoint was called
		switch r.URL.Path {
		case "/success-api-url":
			// Simulate a successful API response with trivia questions
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"response_code": 0,
				"results": [
					{
						"type": "multiple",
						"difficulty": "easy",
						"category": "General Knowledge",
						"question": "What is 2 + 2?",
						"correct_answer": "4",
						"incorrect_answers": ["2", "3", "5"]
					}
				]
			}`))
		case "/error-api-url":
			// Simulate an API response with a non-zero response code (indicating error)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"response_code": 1,
				"results": []
			}`))
		default:
			// Return 404 for unknown paths
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer mockServer.Close()

	// Test case for successful API response
	t.Run("Successful API response", func(t *testing.T) {
		apiURL := mockServer.URL + "/success-api-url"
		questions, err := fetchQuestions(apiURL)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		expectedQuestions := []triviaQuestion{
			{
				Type:             "multiple",
				Difficulty:       "easy",
				Category:         "General Knowledge",
				Question:         "What is 2 + 2?",
				CorrectAnswer:    "4",
				IncorrectAnswers: []string{"2", "3", "5"},
			},
		}

		if len(questions) != 1 || !compareTriviaQuestions(questions[0], expectedQuestions[0]) {
			t.Errorf("Unexpected questions returned:\nGot: %v\nExpected: %v", questions, expectedQuestions)
		}
	})

	// Test case for API error (No Results)
	t.Run("API error: No Results", func(t *testing.T) {
		apiURL := mockServer.URL + "/error-api-url"
		_, err := fetchQuestions(apiURL)
		expectedErrMsg := "opentdb:API error with code:1 ,check docs for more info"
		if err == nil || err.Error() != expectedErrMsg {
			t.Errorf("Unexpected error: %v", err)
		}
	})

	// Test case for API error (HTTP 404)
	t.Run("API error: HTTP 404", func(t *testing.T) {
		apiURL := mockServer.URL + "/unknown-api-url"
		_, err := fetchQuestions(apiURL)
		expectedErrMsg := "API request failed with status: 404 Not Found"
		if err == nil || err.Error() != expectedErrMsg {
			t.Errorf("Unexpected error: %v", err)
		}
	})
}
func TestFetchCategories(t *testing.T) {
	// Create a mock HTTP server to simulate the API response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the URL path
		if r.URL.Path != "/api_category.php" {
			t.Errorf("Unexpected URL path. Got: %s, Expected: /api_category.php", r.URL.Path)
		}

		// Simulate a successful API response with trivia categories
		categories := []question.Category{
			{Id: 9, Name: "General Knowledge"},
			{Id: 10, Name: "Entertainment: Books"},
			{Id: 11, Name: "Entertainment: Film"},
		}
		categoryResp := categoryResponse{TriviaCategories: categories}

		// Encode and write the JSON response
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(categoryResp)
		if err != nil {
			t.Errorf("Failed to encode JSON response: %v", err)
		}
	}))
	defer mockServer.Close()

	// Use the mock server's URL to fetch categories
	url := mockServer.URL + "/api_category.php"
	categories, err := fetchCategories(url)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Define the expected categories based on the mock response
	expectedCategories := []question.Category{
		{Id: 9, Name: "General Knowledge"},
		{Id: 10, Name: "Entertainment: Books"},
		{Id: 11, Name: "Entertainment: Film"},
	}

	// Compare the fetched categories with the expected categories
	if !reflect.DeepEqual(categories, expectedCategories) {
		t.Errorf("Unexpected categories. Got: %v, Expected: %v", categories, expectedCategories)
	}
}
func TestShuffleOptions(t *testing.T) {

	t.Run("Shuffle options", func(t *testing.T) {

		options := []string{"Apple", "Banana", "Cherry", "Date"}

		originalOptions := make([]string, len(options))
		copy(originalOptions, options)

		shuffleOptions(options)

		if len(options) != len(originalOptions) {
			t.Errorf("Length of shuffled options doesn't match original")
		}

		for _, opt := range originalOptions {
			if !contains(options, opt) {
				t.Errorf("Shuffled options are missing element: %s", opt)
			}
		}
	})
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func compareTriviaQuestions(q1, q2 triviaQuestion) bool {
	return q1.Type == q2.Type &&
		q1.Difficulty == q2.Difficulty &&
		q1.Category == q2.Category &&
		q1.Question == q2.Question &&
		q1.CorrectAnswer == q2.CorrectAnswer &&
		sliceEqual(q1.IncorrectAnswers, q2.IncorrectAnswers)
}
func sliceEqual(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

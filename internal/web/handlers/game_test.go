package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/Lafetz/showdown-trivia-game/internal/core/entities"
	"github.com/Lafetz/showdown-trivia-game/internal/core/question"
	"github.com/Lafetz/showdown-trivia-game/internal/web/ws"
	"github.com/PuerkitoBio/goquery"
)

type mockapi struct{}

func (m *mockapi) GetCategories() ([]question.Category, error) {
	return []question.Category{}, nil
}
func (m *mockapi) GetQuestions(amount int, category int) ([]entities.Question, error) {
	return []entities.Question{}, nil
}
func TestHome(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req = req.WithContext(context.WithValue(req.Context(), UsernameKey, "testuser"))

	w := httptest.NewRecorder()

	handler := Home(mockLogger)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, w.Code)
	}

	expectedHTML := `hx-get="/activegames"`
	verifyHtml(t, w, expectedHTML)

}

func TestCreateGet(t *testing.T) {
	req := httptest.NewRequest("GET", "/create", nil)
	w := httptest.NewRecorder()
	mockapi := &mockapi{}
	questionService := question.NewQuestionService(mockapi)
	handler := CreateFormGet(mockLogger, questionService)
	handler.ServeHTTP(w, req)

	// Check response status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, w.Code)
	}

	expectedHTML := `hx-post="/create"`
	verifyHtml(t, w, expectedHTML)

}

func TestActiveGames(t *testing.T) {
	hub := &ws.Hub{} // Mock WebSocket hub
	req := httptest.NewRequest("GET", "/active", nil)
	w := httptest.NewRecorder()

	handler := ActiveGames(hub, mockLogger)
	handler.ServeHTTP(w, req)

	// Check response status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, w.Code)
	}

}

func TestJoin(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handler := Join(mockLogger, "connect:ws://localhost:8080")
	handler.ServeHTTP(w, req)

	// Check response status code
	if w.Code != http.StatusOK {
		t.Errorf("Join handler returned unexpected status code: %d", w.Code)
	}

	expectedHTML := `<div hx-ws="connect:ws://localhost:8080/wsjoin/">`
	verifyHtml(t, w, expectedHTML)

}

type MockQuestionService struct{}

func (m *MockQuestionService) GetCategories() ([]question.Category, error) {

	return []question.Category{}, nil
}
func (m *MockQuestionService) GetQuestions(amount int, category int) ([]entities.Question, error) {
	return []entities.Question{}, nil
}
func TestCreateFormPost(t *testing.T) {

	logger := mockLogger

	testCases := []struct {
		name        string
		formData    url.Values
		expectation string
	}{
		{
			name: "ValidForm_Success",
			formData: url.Values{
				"category": {"1"},  // Mock category ID
				"timer":    {"20"}, // Mock timer value
				"amount":   {"10"}, // Mock amount of questions
			},
			expectation: `id="socket"`,
		},
		{
			name:        "InvalidForm_BadRequest",
			formData:    url.Values{}, // Missing required fields
			expectation: `hx-post="/create"`,
		},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a mock QuestionServiceApi
			mockQuestionService := &MockQuestionService{}

			req := httptest.NewRequest("POST", "/your-endpoint", strings.NewReader(tc.formData.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Set("Content-Length", strconv.Itoa(len(tc.formData.Encode())))

			w := httptest.NewRecorder()

			handler := CreateFormPost(logger, mockQuestionService, "connect:ws://localhost:8080")
			handler.ServeHTTP(w, req)
			verifyHtml(t, w, tc.expectation)
		})
	}
}
func verifyHtml(t *testing.T, w *httptest.ResponseRecorder, expectedHtml string) {
	doc, err := goquery.NewDocumentFromReader(w.Body)
	if err != nil {
		t.Fatalf("Error parsing HTML response: %v", err)
	}

	s, err := doc.Html()
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(s, expectedHtml) {
		t.Errorf("Expected HTML %q not found in rendered output", expectedHtml)
	}
}

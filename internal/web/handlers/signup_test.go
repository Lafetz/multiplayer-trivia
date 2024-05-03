package handlers

import (
	"net/http"
	"net/http/httptest"

	"strings"
	"testing"

	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
	"github.com/PuerkitoBio/goquery"
)

type mockUserService struct{}

func (m *mockUserService) AddUser(user *user.User) (*user.User, error) {

	return user, nil
}
func TestSignupGet(t *testing.T) {
	// Create a mock HTTP request for the SignupGet handler (GET request)
	req := httptest.NewRequest("GET", "/signup", nil)
	w := httptest.NewRecorder()

	// Call the SignupGet handler function directly
	SignupGet().ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("SignupGet handler returned unexpected status code: %d", w.Code)
	}

	doc, err := goquery.NewDocumentFromReader(w.Body)
	if err != nil {
		t.Fatalf("Error parsing HTML response: %v", err)
	}

	// Example validation using goquery selectors
	s, err := doc.Html()
	if err != nil {
		t.Fatal(err)
	}
	expectedHTML := `hx-post="/signup" `
	println(s)
	if !strings.Contains(s, expectedHTML) {
		t.Errorf("Expected HTML %q not found in rendered output", expectedHTML)
	}
}
func TestSignupPost(t *testing.T) {
	userService := &mockUserService{}
	handler := SignupPost(userService)

	t.Run("ValidSignup", func(t *testing.T) {
		formData := strings.NewReader("username=test&email=test@example.com&password=pass123")
		req := httptest.NewRequest("POST", "/signup", formData)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		// Check response status code
		if w.Code != http.StatusCreated {
			t.Errorf("expected status %d; got %d", http.StatusCreated, w.Code)
		}

	})

	t.Run("InvalidSignup", func(t *testing.T) {
		// Simulate invalid form data (missing required fields)
		formData := strings.NewReader("username=&email=&password=")
		req := httptest.NewRequest("POST", "/signup", formData)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		// Check response status code
		if w.Code != http.StatusUnprocessableEntity {
			t.Errorf("expected status %d; got %d", http.StatusUnprocessableEntity, w.Code)
		}

	})

}

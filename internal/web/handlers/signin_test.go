package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/sessions"
	// Import the package containing SigninPost
)

// Mock UserService implementation for testing

func (m *mockUserService) GetUser(email string) (*user.User, error) {
	// Simulate fetching a user
	p, err := hashPassword("pass123456")
	if err != nil {
		return nil, err
	}

	if email == "test@example.com" {
		return &user.User{
			Username: "testuser",
			Email:    "test@example.com",
			Password: p, // This should be a hashed password
		}, nil
	}
	return nil, errors.New("not found") // Simulate user not found or error
}
func TestSigninGetHandler(t *testing.T) {
	// Create a mock HTTP request for the SigninGet handler (GET request)
	req := httptest.NewRequest("GET", "/signin", nil)
	w := httptest.NewRecorder()

	SigninGet(mockLogger).ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("SigninGet handler returned unexpected status code: %d", w.Code)
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
	expectedHTML := `hx-post="/signin"`
	println(s)
	if !strings.Contains(s, expectedHTML) {
		t.Errorf("Expected HTML %q not found in rendered output", expectedHTML)
	}

}
func TestSigninPost(t *testing.T) {
	// Initialize the SigninPost handler with a mock UserService and mock CookieStore
	userService := &mockUserService{}
	hashKey := "your-generated-hash-key"
	blockKey := "your-generated-block-key"
	store := sessions.NewCookieStore([]byte(hashKey), []byte(blockKey)) // Create a mock cookie store

	handler := SigninPost(userService, store, mockLogger)

	t.Run("ValidSignin", func(t *testing.T) {
		formData := strings.NewReader("email=test@example.com&password=pass123456")
		req := httptest.NewRequest("POST", "/signin", formData)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		// Check response status code
		if w.Code != http.StatusOK {
			t.Errorf("expected status %d; got %d", http.StatusOK, w.Code)
		}

	})

	t.Run("InvalidForm", func(t *testing.T) {

		formData := strings.NewReader("email=&password=")
		req := httptest.NewRequest("POST", "/signin", formData)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		if w.Code != http.StatusUnprocessableEntity {
			t.Errorf("expected status %d; got %d", http.StatusUnprocessableEntity, w.Code)
		}

	})

	t.Run("UserNotFound", func(t *testing.T) {
		// Simulate user not found in the database
		formData := strings.NewReader("email=nonexistent@example.com&password=pass123456")
		req := httptest.NewRequest("POST", "/signin", formData)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		// Check response status code
		if w.Code != http.StatusUnauthorized {
			t.Errorf("expected status %d; got %d", http.StatusUnauthorized, w.Code)
		}

	})

	t.Run("IncorrectPassword", func(t *testing.T) {
		// Simulate incorrect password
		formData := strings.NewReader("email=test@example.com&password=wrongpassword")
		req := httptest.NewRequest("POST", "/signin", formData)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		// Check response status code
		if w.Code != http.StatusUnauthorized {
			t.Errorf("expected status %d; got %d", http.StatusUnauthorized, w.Code)
		}

	})
}

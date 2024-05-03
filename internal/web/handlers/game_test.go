package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Lafetz/showdown-trivia-game/internal/ws"
	"github.com/PuerkitoBio/goquery"
)

func TestHome(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req = req.WithContext(context.WithValue(req.Context(), "username", "testuser"))

	w := httptest.NewRecorder()

	handler := Home()
	handler.ServeHTTP(w, req)

	// Check response status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, w.Code)
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
	expectedHTML := `<section id="create-game" class="flex flex-col items-center justify-center gap-7">`
	println(s)
	if !strings.Contains(s, expectedHTML) {
		t.Errorf("Expected HTML %q not found in rendered output", expectedHTML)
	}
}

func TestCreateGet(t *testing.T) {
	req := httptest.NewRequest("GET", "/create", nil)
	w := httptest.NewRecorder()

	handler := CreateGet()
	handler.ServeHTTP(w, req)

	// Check response status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, w.Code)
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
	expectedHTML := `<div id="socket" hx-ws="connect:ws://localhost:8080/wscreate">`

	if !strings.Contains(s, expectedHTML) {
		t.Errorf("Expected HTML %q not found in rendered output", expectedHTML)
	}
}

func TestActiveGames(t *testing.T) {
	hub := &ws.Hub{} // Mock WebSocket hub
	req := httptest.NewRequest("GET", "/active", nil)
	w := httptest.NewRecorder()

	handler := ActiveGames(hub)
	handler.ServeHTTP(w, req)

	// Check response status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, w.Code)
	}

}

func TestJoin(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handler := Join()
	handler.ServeHTTP(w, req)

	// Check response status code
	if w.Code != http.StatusOK {
		t.Errorf("Join handler returned unexpected status code: %d", w.Code)
	}

	// Validate rendered HTML using goquery
	doc, err := goquery.NewDocumentFromReader(w.Body)
	if err != nil {
		t.Fatalf("Error parsing HTML response: %v", err)
	}

	// Example validation using goquery selectors
	s, err := doc.Html()
	if err != nil {
		t.Fatal(err)
	}
	expectedHTML := `<div hx-ws="connect:ws://localhost:8080/wsjoin/">`

	if !strings.Contains(s, expectedHTML) {
		t.Errorf("Expected HTML %q not found in rendered output", expectedHTML)
	}

}

package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
)

func TestSignout(t *testing.T) {
	// Initialize a new cookie store for session management
	store := sessions.NewCookieStore([]byte("your-secret-key"))

	// Create a new HTTP request (simulating a sign-out request)
	req, err := http.NewRequest("GET", "/signout", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP test recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a new Signout handler function with a mock logger
	handler := Signout(mockLogger, store)

	// Serve the HTTP request using the Signout handler function
	handler.ServeHTTP(rr, req)

	// Check the HTTP response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verify that the session has been updated correctly
	session, err := store.Get(req, "user-session")
	if err != nil {
		t.Errorf("failed to get session: %v", err)
	}

	// Check the value of 'authenticated' in the session
	if authenticated, ok := session.Values["authenticated"].(bool); !ok || authenticated {
		t.Errorf("session 'authenticated' value not updated correctly: got %v want false", authenticated)
	}

	// Check the HX-Redirect header value
	expectedRedirect := "/signin"
	if redirect := rr.Header().Get("HX-Redirect"); redirect != expectedRedirect {
		t.Errorf("unexpected redirect header: got %s want %s", redirect, expectedRedirect)
	}
}

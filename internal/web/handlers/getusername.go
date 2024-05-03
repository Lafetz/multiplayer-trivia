package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

func getSessionUsername(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore) (string, error) {

	session, err := store.Get(r, "user-session")
	if err != nil {
		http.Error(w, "Session not found", http.StatusUnauthorized)
		return "", err
	}

	authenticated, ok := session.Values["authenticated"].(bool)
	if !ok || !authenticated {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return "", fmt.Errorf("user not authenticated")
	}

	username, ok := session.Values["username"].(string)
	if !ok {
		http.Error(w, "Username not found in session", http.StatusInternalServerError)
		return "", fmt.Errorf("username not found in session")
	}

	return username, nil
}

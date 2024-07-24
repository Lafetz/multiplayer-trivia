package ws

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Lafetz/showdown-trivia-game/internal/core/entities"

	"github.com/gorilla/websocket"
)

func TestRoomAddClientRemoveClient(t *testing.T) {
	// Mock WebSocket server for testing
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Upgrade the HTTP server to a WebSocket connection
		conn, err := (&websocket.Upgrader{}).Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("failed to upgrade connection to WebSocket: %v", err)
		}

		// Create a new room for testing
		roomID := "testRoom"
		questions := []entities.Question{
			{Question: "What is 2+2?", Options: []string{"3", "4", "5", "6"}, CorrectAnswer: "4"},
			{Question: "What is the capital of France?", Options: []string{"London", "Berlin", "Paris", "Rome"}, CorrectAnswer: "Paris"},
		}
		owner := "test"
		room := NewRoom(roomID, owner, 2, questions, &Hub{})

		// Add a client to the room
		client := NewClient(conn, room, "1")
		room.addClient(client)

		// Ensure the client is added to the room's client list
		if _, ok := room.clients[client]; !ok {
			t.Error("failed to add client to room")
		}

		// Check if the room sends updated player list after client addition
		expectedPlayerList := []string{client.Username}
		if !equalSlices(room.getUsers(), expectedPlayerList) {
			t.Errorf("expected player list %v, got %v", expectedPlayerList, room.getUsers())
		}

		// Simulate sending a message from the client
		clientMessage := "test message"
		client.egress <- []byte(clientMessage)

		// Check if the message was sent to all clients in the room
		select {
		case msg := <-client.egress:
			if string(msg) != clientMessage {
				t.Errorf("expected message '%s', got '%s'", clientMessage, string(msg))
			}
		default:
			t.Error("failed to send message to clients in the room")
		}

		// Remove the client from the room
		err = room.removeClient(client)
		if err != nil {
			t.Fatal(err)
		}
		// Ensure the client is removed from the room's client list
		if _, ok := room.clients[client]; ok {
			t.Error("failed to remove client from room")
		}

		// Check if the room sends updated player list after client removal
		if len(room.getUsers()) != 0 {
			t.Errorf("expected no players in room, got %v", room.getUsers())
		}
	}))

	defer server.Close()

	// Connect to the mock WebSocket server
	url := "ws" + strings.TrimPrefix(server.URL, "http")
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		t.Fatalf("failed to connect to WebSocket server: %v", err)
	}
	defer conn.Close()
}

// Helper function to check equality of string slices
func equalSlices(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

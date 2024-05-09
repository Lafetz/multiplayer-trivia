package ws

import (
	"errors"
	"log"
	"reflect"
	"testing"

	"github.com/Lafetz/showdown-trivia-game/internal/core/entities"
	webentities "github.com/Lafetz/showdown-trivia-game/internal/web/entity"
)

func TestHubGetRoom(t *testing.T) {
	hub := NewHub(log.Default())

	// Create a test room
	roomId := "testRoom"
	questions := []entities.Question{
		{Question: "What is 2+2?", Options: []string{"3", "4", "5", "6"}, CorrectAnswer: "4"},
		{Question: "What is the capital of France?", Options: []string{"London", "Berlin", "Paris", "Rome"}, CorrectAnswer: "Paris"},
	}
	owner := "test"
	room := NewRoom(roomId, owner, 2, questions, hub)
	hub.addRoom(room)

	// Test getting an existing room
	resultRoom, err := hub.getRoom(roomId)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resultRoom.Id != room.Id {
		t.Errorf("expected room ID %s, got %s", room.Id, resultRoom.Id)
	}
	if resultRoom.owner != room.owner {
		t.Errorf("expected room owner %s, got %s", room.owner, resultRoom.owner)
	}
	// if !reflect.DeepEqual(resultRoom, room) {
	// 	t.Errorf("expected room %+v, got %+v", room, resultRoom)
	// }

	// Test getting a non-existing room
	nonExistingRoomId := "nonExistingRoom"
	_, err = hub.getRoom(nonExistingRoomId)
	if err == nil {
		t.Error("expected error for non-existing room, but got nil")
	}
	if !errors.Is(err, ErrRoomNotExist) {
		t.Errorf("expected 'room doesn't exist' error, got %v", err)
	}
}

func TestHubAddRoomRemoveRoom(t *testing.T) {
	hub := NewHub(log.Default())

	// Create a test room
	questions := []entities.Question{
		{Question: "What is 2+2?", Options: []string{"3", "4", "5", "6"}, CorrectAnswer: "4"},
		{Question: "What is the capital of France?", Options: []string{"London", "Berlin", "Paris", "Rome"}, CorrectAnswer: "Paris"},
	}
	owner := "test"
	roomId := "testRoom"
	room := NewRoom(roomId, owner, 2, questions, hub)

	// Test adding a room
	hub.addRoom(room)
	if _, ok := hub.rooms[roomId]; !ok {
		t.Errorf("room was not added to hub")
	}

	// Test listing rooms
	rooms := hub.ListRooms()
	if len(rooms) != 1 {
		t.Errorf("expected one room in hub, got %d", len(rooms))
	}

	// Test removing a room
	hub.removeRoom(room)
	if _, ok := hub.rooms[roomId]; ok {
		t.Errorf("room was not removed from hub")
	}
}

func TestHubListRooms(t *testing.T) {
	hub := NewHub(log.Default())

	questions := []entities.Question{
		{Question: "What is 2+2?", Options: []string{"3", "4", "5", "6"}, CorrectAnswer: "4"},
		{Question: "What is the capital of France?", Options: []string{"London", "Berlin", "Paris", "Rome"}, CorrectAnswer: "Paris"},
	}
	owner := "test"
	// Create test rooms
	room1 := NewRoom("room1", owner, 2, questions, hub)
	room2 := NewRoom("room2", owner, 2, questions, hub)

	// Add rooms to hub
	hub.addRoom(room1)
	hub.addRoom(room2)

	// Test listing rooms
	expectedRooms := []webentities.RoomData{
		{Owner: room1.owner, Id: room1.Id, Players: room1.getUsers()},
		{Owner: room2.owner, Id: room2.Id, Players: room2.getUsers()},
	}
	actualRooms := hub.ListRooms()

	if len(actualRooms) != len(expectedRooms) {
		t.Errorf("number of listed rooms does not match expected")
	}

	for i, expected := range expectedRooms {
		if !reflect.DeepEqual(actualRooms[i], expected) {
			//fmt.Printf("%s", actualRooms[i])
			t.Errorf("listed room at index %d does not match expected", i)
		}
	}
}

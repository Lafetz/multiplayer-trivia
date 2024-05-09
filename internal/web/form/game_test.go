package form

import "testing"

func TestValid(t *testing.T) {
	tests := []struct {
		name     string
		game     NewGame
		expected bool
	}{
		{
			name: "Valid input",
			game: NewGame{
				Category: "10",
				Timer:    "5",
				Amount:   "30",
			},
			expected: true,
		},
		{
			name: "Invalid category (not a number)",
			game: NewGame{
				Category: "abc",
				Timer:    "5",
				Amount:   "30",
			},
			expected: false,
		},
		{
			name: "Invalid timer (out of range)",
			game: NewGame{
				Category: "10",
				Timer:    "25",
				Amount:   "30",
			},
			expected: false,
		},
		{
			name: "Invalid amount (out of range)",
			game: NewGame{
				Category: "10",
				Timer:    "5",
				Amount:   "100",
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.game.Valid()

			if actual != tc.expected {
				t.Errorf("Expected Valid() to be %v, but got %v", tc.expected, actual)
			}
		})
	}

}

func TestValidErrorMessages(t *testing.T) {
	invalidGame := NewGame{
		Category: "abc",
		Timer:    "23",
		Amount:   "100",
	}
	invalidGame.Valid()

	expectedErrors := map[string]string{
		"category": "needs to be number",
		"timer":    "timer can't be less than 2 or greater than 20",
		"amount":   "number of questions can't be less than 1 or greater than 50",
	}

	for field, expectedMsg := range expectedErrors {
		actualMsg, ok := invalidGame.Errors[field]
		if !ok {
			t.Errorf("Expected error message for %s field, but not found", field)
		} else if actualMsg != expectedMsg {
			t.Errorf("Expected error message for %s field to be %q, but got %q", field, expectedMsg, actualMsg)
		}
	}
}

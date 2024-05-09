package form

import "testing"

func TestSigninUserValid(t *testing.T) {
	// Test case: Valid inputs
	validUser := &SigninUser{
		Email:    "john@example.com",
		Password: "strongpassword",
	}
	if !validUser.Valid() {
		t.Errorf("Expected valid user, got invalid")
	}

	// Test case: Missing email
	missingEmail := &SigninUser{
		Email:    "", // Empty email
		Password: "password123",
	}
	if missingEmail.Valid() {
		t.Errorf("Expected invalid user (missing email), got valid")
	}

	// Test case: Missing password
	missingPassword := &SigninUser{
		Email:    "joe@example.com",
		Password: "", // Empty password
	}
	if missingPassword.Valid() {
		t.Errorf("Expected invalid user (missing password), got valid")
	}

	// Test case: Weak password
	weakPassword := &SigninUser{
		Email:    "jane@example.com",
		Password: "weak", // Too short password
	}
	if weakPassword.Valid() {
		t.Errorf("Expected invalid user (weak password), got valid")
	}
}
func TestSignupUserValid(t *testing.T) {
	// Test case: Valid inputs
	validUser := &SignupUser{
		Username: "john_doe",
		Email:    "john@example.com",
		Password: "strongpassword",
	}
	if !validUser.Valid() {
		t.Errorf("Expected valid user, got invalid")
	}

	// Test case: Missing username
	missingUsername := &SignupUser{
		Username: "", // Empty username
		Email:    "jane@example.com",
		Password: "password123",
	}
	if missingUsername.Valid() {
		t.Errorf("Expected invalid user (missing username), got valid")
	}
	// Test case: invalid username
	invalidUsername := &SignupUser{
		Username: "@hello", // Empty username
		Email:    "jane@example.com",
		Password: "password123",
	}
	if invalidUsername.Valid() {
		t.Errorf("Expected invalid user (missing username), got valid")
	}
	// Test case: Invalid email format
	invalidEmail := &SignupUser{
		Username: "jane_doe",
		Email:    "invalid_email",
		Password: "password123",
	}
	if invalidEmail.Valid() {
		t.Errorf("Expected invalid user (invalid email), got valid")
	}

	// Test case: Weak password
	weakPassword := &SignupUser{
		Username: "joe_smith",
		Email:    "joe@example.com",
		Password: "weak", // Too short password
	}
	if weakPassword.Valid() {
		t.Errorf("Expected invalid user (weak password), got valid")
	}
}

package handlers

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
	// Import the package containing hashPassword and matchPassword
)

func TestHashPassword(t *testing.T) {
	plaintextPassword := "testpassword"

	hash, err := hashPassword(plaintextPassword)
	if err != nil {
		t.Errorf("unexpected error hashing password: %v", err)
	}

	// Validate the generated hash
	err = bcrypt.CompareHashAndPassword(hash, []byte(plaintextPassword))
	if err != nil {
		t.Errorf("generated hash is invalid: %v", err)
	}
}

func TestMatchPassword(t *testing.T) {
	plaintextPassword := "testpassword"
	validHash, _ := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)

	t.Run("ValidPassword", func(t *testing.T) {
		err := matchPassword(plaintextPassword, validHash)
		if err != nil {
			t.Errorf("expected password to match: %v", err)
		}
	})

	t.Run("InvalidPassword", func(t *testing.T) {
		invalidPassword := "wrongpassword"
		err := matchPassword(invalidPassword, validHash)
		expectedErr := ErrInvalidPassword
		if err == nil || err != expectedErr {
			t.Errorf("expected %v error for invalid password", expectedErr)
		}
	})

	t.Run("InvalidHash", func(t *testing.T) {
		invalidHash := []byte("invalidhash")
		err := matchPassword(plaintextPassword, invalidHash)
		if err == nil {
			t.Errorf("expected error with invalid hash")
		}
	})
}

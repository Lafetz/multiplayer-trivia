package handlers

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidPassword = errors.New("incorrect password")
)

func hashPassword(plaintextPassword string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return nil, err
	}

	return hash, nil
}
func matchPassword(plaintextPassword string, hash []byte) error {

	err := bcrypt.CompareHashAndPassword(hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return ErrInvalidPassword
		default:
			return err
		}
	}

	return err
}

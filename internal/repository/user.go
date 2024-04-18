package repository

import (
	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
	"github.com/google/uuid"
)

func (store *Store) GetUser(username string) (*user.User, error) {
	return user.NewUser(username, "email!gmail.com", []byte("sss")), nil
}
func (store *Store) AddUser(user *user.User) (*user.User, error) {
	return user, nil
}

func (store *Store) DeleteUser(id uuid.UUID) error {
	return nil
}

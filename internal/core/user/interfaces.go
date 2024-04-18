package user

import "github.com/google/uuid"

type UserServiceApi interface {
	GetUser(string) (*User, error)
	AddUser(*User) (*User, error)
	DeleteUser(uuid.UUID) error
}
type repository interface {
	GetUser(username string) (*User, error)
	AddUser(*User) (*User, error)
	DeleteUser(uuid.UUID) error
}

package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	Username  string
	Email     string
	Password  []byte
	CreatedAt time.Time
}

func NewUser(username string, email string, password []byte) *User {
	user := &User{
		Id:        uuid.New(),
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}

	return user
}

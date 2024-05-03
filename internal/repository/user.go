package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type userMongo struct {
	Id        uuid.UUID `bson:"_id,omitempty"`
	Username  string    `bson:"username"`
	Email     string    `bson:"email"`
	Password  []byte    `bson:"password"`
	CreatedAt time.Time `bson:"createdAt"`
}

// domain converts userMongo to user.User
func (u *userMongo) domain() *user.User {
	return &user.User{
		Id:        u.Id,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
	}
}

func newUserMongo(u *user.User) *userMongo {
	return &userMongo{
		Id:        u.Id,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
	}
}

func (store *Store) AddUser(user *user.User) (*user.User, error) {
	u := newUserMongo(user)
	r, err := store.users.InsertOne(context.Background(), u)
	fmt.Printf("over here %s", r)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (store *Store) GetUser(email string) (*user.User, error) {
	var user userMongo
	err := store.users.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user.domain(), nil

}

func (store *Store) DeleteUser(Id uuid.UUID) error {
	return nil
}

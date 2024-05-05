package repository

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/Lafetz/showdown-trivia-game/internal/core/user"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userMongo struct {
	Id        uuid.UUID `bson:"_id,omitempty"`
	Username  string    `bson:"username" unique:"true"`
	Email     string    `bson:"email" unique:"true"`
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

func (store *Store) AddUser(userData *user.User) (*user.User, error) {
	u := newUserMongo(userData)
	_, err := store.users.InsertOne(context.Background(), u)
	if err != nil {
		if mongoErr, ok := err.(mongo.WriteException); ok {

			for _, writeErr := range mongoErr.WriteErrors {
				if writeErr.Code == 11000 {

					key := extractDuplicateKey(writeErr.Message)
					switch key {
					case "email":
						return nil, user.ErrEmailUnique
					case "username":
						return nil, user.ErrUsernameUnique
					default:
						return nil, err

					}
				}
			}
		} else {
			return nil, err
		}
		return userData, nil
	}
	return userData, nil
}
func (store *Store) GetUser(email string) (*user.User, error) {
	var userData userMongo
	err := store.users.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&userData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, user.ErrUserNotFound
		}
		return nil, err
	}

	return userData.domain(), nil

}

func (store *Store) DeleteUser(Id uuid.UUID) error {
	return nil
}
func extractDuplicateKey(errorMessage string) string {

	pattern := `index: ([a-zA-Z0-9_]+)_\d+ dup key`

	re := regexp.MustCompile(pattern)

	match := re.FindStringSubmatch(errorMessage)
	if len(match) < 2 {
		return ""
	}

	indexName := match[1]

	parts := strings.Split(indexName, "_")
	if len(parts) > 0 {
		fieldName := parts[0]
		return fieldName
	}

	return ""
}

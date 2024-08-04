package repository

import (
	"context"
	"log/slog"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	users *mongo.Collection
}

func NewDb(url string, logger *slog.Logger) (*mongo.Client, func(), error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		return nil, nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	return client, func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := client.Disconnect(ctx)
		if err != nil {
			logger.Error(err.Error())
		}

	}, nil

}

func NewStore(client *mongo.Client) (*Store, error) {

	users := client.Database("trivia").Collection("users")
	err := createUniqueIndex(context.Background(), users, "email")
	if err != nil {
		return nil, err
	}
	err = createUniqueIndex(context.Background(), users, "username")
	if err != nil {
		return nil, err
	}

	return &Store{
		users: users,
	}, nil
}
func createUniqueIndex(ctx context.Context, collection *mongo.Collection, fieldName string) error {
	// Create a unique index model for the specified field
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: fieldName, Value: 1}}, // Specify the field and its sorting order (1 for ascending)
		Options: options.Index().SetUnique(true),    // Set the unique option to true for enforcing uniqueness
	} //

	// Create the unique index
	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}

	return nil
}

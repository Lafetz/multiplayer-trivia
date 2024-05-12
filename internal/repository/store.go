package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	users *mongo.Collection
}

func NewDb(url string) *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal("db connection falled", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("db connection falled", err)
	}
	return client
}

func NewStore(client *mongo.Client) *Store {

	users := client.Database("trivia").Collection("users")
	err := createUniqueIndex(context.Background(), users, "email")
	if err != nil {
		log.Fatal("Failed to create unique index:", err)
	}
	err = createUniqueIndex(context.Background(), users, "username")
	if err != nil {
		log.Fatal("Failed to create unique index:", err)
	}
	println("database connected.....")
	return &Store{
		users: users,
	}
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
	fmt.Printf("Unique index created on field '%s'\n", fieldName)
	return nil
}

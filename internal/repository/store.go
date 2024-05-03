package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	users *mongo.Collection
}

func NewStore() *Store {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://admin:admin11@localhost/inventory?authSource=admin"))
	if err != nil {
		log.Fatal("db connection falled", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("db connection falled", err)
	}
	users := client.Database("trivia").Collection("users")
	println("database connected.....")
	return &Store{
		users: users,
	}
}

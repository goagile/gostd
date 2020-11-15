package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	uri := "mongodb://127.0.0.1:27017"

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Millisecond)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("mongo.NewClient ", err)
	}
	defer client.Disconnect(ctx)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("client.Connect ", err)
	}

	r, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal("ListDatabaseNames ", err)
	}
	fmt.Println(r)
}

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Movie struct {
	ID 			primitive.ObjectID	`bson:"_id,omitempty"`
	Title		string				`bson:"title,omitempty"`
	Plot		string				`bson:"author,omitempty"`
	Year		uint16				`bson:"year,omitempty"`
	Released	time.Time			`bson:"released,omitempty"`
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer client.Disconnect(ctx)

	database := client.Database("mflix")
	moviesCollection := database.Collection("movies")

	wakanga := Movie{ Title: "Wakanga", Plot: "Black people", Year: 2020, Released: time.Now() }

	result, err := moviesCollection.InsertOne(ctx, wakanga)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println("Result", result)
}
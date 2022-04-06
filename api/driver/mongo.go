package driver

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func MongoDB() {
	/*
	   Connect to my cluster
	*/
	fmt.Println("DB 가동 시작")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	/*
	   List databases
	*/
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	// Interact with data
	type User struct {
		Name string `bson:"name,omitempty"`
		Age  int    `bson:"age,omitempty"`
	}

	/*
		Get my collection instance
	*/
	collection := client.Database("clone").Collection("users")

	/*
		Insert documents
	*/
	docs := []interface{}{
		bson.D{{"name", "a"}, {"age", 1}},
		bson.D{{"name", "b"}, {"age", 2}},
		bson.D{{"name", "c"}, {"age", 3}},
	}

	res, insertErr := collection.InsertMany(ctx, docs)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	fmt.Println(res)
	/*
		Iterate a cursor
	*/
	cur, currErr := collection.Find(ctx, bson.D{})

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	var posts []User
	if err = cur.All(ctx, &posts); err != nil {
		panic(err)
	}
	fmt.Println(posts)
}

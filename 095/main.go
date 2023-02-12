// Aggregations
// In this exercises, we are going to learn about BSON, a binary serialization format (like JSON) which is used to marshall and unmarshall data ad make remote calls in mongoDB
// First, we need to import the `go.mongodb.org/mongo-driver/bson` library.
package main

import (
	"context"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct{
	Name string
	Age  int
}

func main (){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	usersCollection := client.Database("TestCluster").Collection("users")
	log.Println(usersCollection.Name())
	
	// We are going to create a group stage, this will have the result for the average age for each named user!.
	// Create a variable named groupStage 
	// we will assign it a bson.D object with the following information:
	/*
		{"$group", bson.D{
			{"_id", "$name"},
			{"average_price", bson.D{{"$avg", "$age"}}},
			{"numTimes", bson.D{{"$sum", 1}}},
		}},
	*/
	
	
	// use the Aggregate() function with the second argument as mongo.Pipeline{groupStage}:
	cursor, err := 
	if err != nil {
		panic(err)
	}
	
	// display the results
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
	    panic(err)
	}

	for _, result := range results {
	    log.Printf("Average age of users named like %v user options: %v \n", result["_id"], result["average_price"])
	    log.Printf("Number of %v tea options: %v \n\n", result["_id"], result["numTimes"])
	}	
}
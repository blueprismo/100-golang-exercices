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

func main (){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
	}

	// Create a variable named "uri" containing your MONGODB_URI string connection.
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

	// Create a collection called "users" into your Database.
	// To do so, use the client.Database() call, and call the .Collection() method to create a collection if it doesn't exist!
	// We will store the returned value into a variable named "usersCollection"
	client.Database("TestCluster").CreateCollection(context.TODO(), "users", options.CreateCollection().SetMaxDocuments(100000))
	usersCollection := client.Database("TestCluster").Collection("users")
	// Here, we will return the name of the collection, to check everything went allright :)
	log.Println(usersCollection.Name())
	log.Println("You got connected!")
	
}
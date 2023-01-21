// Bulk insert
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
	
	// Let's add a bunch of users with the InsertMany() method
	// Create a userList of minimum 5 users, and one of them must be called "Rose"
	userList := []interface{}{
		
	}
	res, err := 
	if err != nil {
		panic(err)
	}
	// And now let's find one user named "John" in the database, let's create a search filter with has the "name" = "John"
	filter := bson.D{{"name","Rose"}}
	// Let's create a return variable of type bson.D
	var result bson.D 
	// Let's use the FindOne() method and reference the return value in the result variable created above
	err = usersCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		panic(err)
	}

	log.Printf("inserted documents with IDs %v\n", res.InsertedIDs)
	
}
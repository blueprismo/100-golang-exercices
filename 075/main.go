package main

// In this exercise we will Stablish a connection with mongoDB - Atlas
// MongoDB is a non relational kind of db, and it's saas has a free tier. Find out more here.
// https://www.mongodb.com/docs/atlas/tutorial/deploy-free-tier-cluster/
// You can also use a docker container locally to connect, use whatever option suits you best.

// 1- You need to install these two libraries through the following commands, one is for the mongo driver and the other for a dotenv file
// go get go.mongodb.org/mongo-driver/mongo
// go get github.com/joho/godotenv
// 2- Once installed, create a file named ".env" containing your connection string. Use that env file to connect to the db.
//    This .env file, will have one variable called "MONGODB_URI=" with the value of your mongodb connection string.
import (
	"context"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main (){
	// Load the godotenv library
	err := 

	if err != nil {
		log.Fatal("Could not load .env file")
	}

	// Thanks to the dotenv file, you can get the MONGODB_URI from the environment!
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	log.Println("You got connected!")
}
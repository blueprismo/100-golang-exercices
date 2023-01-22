// REST API - CRUD operations
// In this exercises, we are going to create an API with the gofiber library.
// https://github.com/gofiber/fiber
// You need to go to the root of this repository and run `go get -u github.com/gofiber/fiber/v2`
package main

import (
	"context"
	"log"
	"os"
	"encoding/json"
	"net/http"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gofiber/fiber/v2"
)

// IMPORTANT FOR THIS EXERCISE, READ ALL THE FILE BEFORE BEGINNING!!! 
// Our goal is to create a new endpoint (app.Get(/users/:id)), so when we try to find /users/John we find one, but when we try to find /users/john it doesn't
// (case sensitivity for this case :) )

type User struct {
	Age  int    `bson:"age"`
	Name string `bson:"name"`
}

// Create a UserResponse type, it will have 3 elements:
// 1- Status, integer
// 2- Message, string
// 3- Data, a *fiber.Map element containing our actual payload. *fiber.Map is a shortcut for map[string]interface{}, useful for JSON returns.
type UserResponse struct {
    Status  int        `json:"status"`
    Message string     `json:"message"`
    Data    *fiber.Map `json:"data"`
}

func initDB() *mongo.Database {
	// MongoDB connection setup
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
	return client.Database("TestCluster")
}

func getUser(c *fiber.Ctx) error {
	// Connect to the DB
	collection := initDB().Collection("users")

	// Get user by name, first extract the name from the URI:
	name :=
	
	// And now let's find one user named "John" in the database, let's create a search filter with has the "name" = parameter we set before
	filter := 
	// Let's create a user variable of type User to reference it with the FindOne function later
	var 
	// Let's use the FindOne() function and reference the return value in the result variable created above!
	err := 
	// Catch the error, if the above query was unsuccesful, return a HTTP500 status message with the err.Error() as the data payload, and "error" as message.
	if err != nil {
        
    }
	
	/*
	// DEBUGGING
	output, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		panic(err)
	}
	log.Printf("%s\n", output)
	*/

	// if everything went alright, return the c.Status(200).JSON(UserResponse{}) with the userresponse status as 200, message as success and the data our user!
	
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	// We will allocate the /user/:name uri, and we will pass the name as queryParam (/user/john)
	
}

// main 
func main (){
	app := fiber.New()
	// Add the routes to the app we created above (right now we only have 2 routes)
	setupRoutes(app)
	// Listen to port 3000
    app.Listen(":3000")
}
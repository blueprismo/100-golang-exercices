// REST API - CRUD operations
// In this exercises, we are going to create an API with the gofiber library.
// https://github.com/gofiber/fiber
// You need to go to the root of this repository and run `go get -u github.com/gofiber/fiber/v2`
package main

import (
	"context"
	"log"
	"os"
	"net/http"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Age  int    `json:"age,string" bson:"age"`
	Name string `json:"name"       bson:"name"`
}

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
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	return client.Database("TestCluster")
}

func getUser(c *fiber.Ctx) error {
	// Connect to the DB
	collection := initDB().Collection("users")
	log.Println(collection.Name())
	// Get user by name, first extract the name from the URI:
	name := c.Params("name")

	// what do we have 
	// And now let's find one user named "John" in the database, let's create a search filter with has the "name" = "John"
	filter := bson.D{{"name",name}}
	// Let's create a return variable of type bson.D
	var user User
	// Let's use the FindOne() method and reference the return value in the result variable created above
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
        return c.Status(500).JSON(UserResponse{Status: 500, Message: "error, not found or unable to find", Data: &fiber.Map{"data": err.Error()}})
    }
	
	return c.Status(http.StatusOK).JSON(UserResponse{Status: 200, Message: "success", Data: &fiber.Map{"data": user}})
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	// We will allocate the /user/:name uri, and we will pass the name as queryParam (/user/john)
	app.Get("/user/:name", getUser)
}

// main 
func main (){
	app := fiber.New()
	// Add the routes to the app we created above (right now we only have 1 route)
	setupRoutes(app)
	// Listen to port 3000
    app.Listen(":3000")
}
// REST API - CRUD operations
// In this exercises, we are going to create an API with the gofiber library.
// https://github.com/gofiber/fiber
// You need to go to the root of this repository and run `go get -u github.com/gofiber/fiber/v2`
package main

import (
	"context"
	"log"
	"os"
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
	db := initDB()
	// Insert the user into the MongoDB collection
	collection := db.Collection("users")
	log.Println(collection.Name())
	// Get user by name, first extract the name from the URI:
	//name, err := primitive.ObjectIDFromHex(c.Params("name"))
	name := c.Params("name")
    /*if err != nil {
        c.Status(400).SendString(err.Error())
        return err
    }*/

	//
	var user User
    err := collection.FindOne(context.Background(), bson.M{"name": name}).Decode(&user)
    if err != nil {
        c.Status(404).SendString(err.Error())
		return err
        
    }
    return c.JSON(user)
	//var user User
	//err = collection.FindOne(context.TODO(),c.Params("name"))
	/*err := collection.FindOne(context.TODO(), bson.M{"name": bson.ObjectIdHex(c.Params("name"))}).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user*/
}

// Create a simple function to return some string to the fiber context
// this function signature will be [func helloWorld(c *fiber.Ctx)]
// inside you will call the Send() method in that context, with the string "Hello World!" as it's only argument.
func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

// Create a function setUP routes with a signature like func setupRoutes(a *fiber.App)
// inside this function use a.GET to the root path, and the helloWorld function we created above
func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	// We will allocate the /user/:name uri, and we will pass the name as queryParam (/user/john)
	app.Get("/user/:name", getUser)
}

// main 
func main (){
	// Fiber 
	// Create a new App with fiber.New()
	app := fiber.New()
	//db := initDB()
	//users_collection := db.Collection("users")
	// Add the routes to the app we created above (right now we only have 1 route)
	setupRoutes(app)
	// Listen to port 3000
    app.Listen(":3000")
}
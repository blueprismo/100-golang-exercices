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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Age  int    `json:"age,string"`
	Name string `json:"name"`
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
	return client.Database("TestCluster")
}

func createUser(c *fiber.Ctx) error {
	// Connect to the DB
	collection := initDB().Collection("users")
	var user User
	err := c.BodyParser(&user);
	if err != nil {
		return c.Status(503).JSON(UserResponse{Status: 503, Message: "error: Incorrect user payload", Data: &fiber.Map{"data": err.Error()}})
	}
	newUser := User{
		Name: user.Name,
		Age : user.Age,
	}
	res, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}
	return c.Status(200).JSON(UserResponse{Status: 200, Message: "success", Data: &fiber.Map{"data": res}})
}

func getUser(c *fiber.Ctx) error {
	// Connect to the DB
	collection := initDB().Collection("users")
	name := c.Params("name")
	filter := bson.D{{"name",name}}
	var user User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
        return c.Status(500).JSON(UserResponse{Status: 500, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }

	return c.Status(200).JSON(UserResponse{Status: 200, Message: "success", Data: &fiber.Map{"data": user}})
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func updateUser(c *fiber.Ctx) error {
	// Connect to the DB
	collection := initDB().Collection("users")
	var user User
	err := c.BodyParser(&user);
	if err != nil {
		return c.Status(503).JSON(UserResponse{Status: 503, Message: "error: Incorrect user payload", Data: &fiber.Map{"data": err.Error()}})
	}
	update := bson.M{"name": user.Name, "age": user.Age}
	filter := bson.D{{"name",c.Params("name")}}
	result, err := collection.UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
        return c.Status(500).JSON(UserResponse{Status: 500, Message: "error: Not found or using capped collection (cannot update a capped collection)", Data: &fiber.Map{"data": err.Error()}})
    }
	// return http 200 when all is fine
	return c.Status(200).JSON(UserResponse{Status: 200, Message: "success", Data: &fiber.Map{"data": result}})
}

func deleteUser(c *fiber.Ctx) error {
	// Connect to the DB
	collection := initDB().Collection("users")
	// Here we only need the filter as the parameter name.
	// in prod, you should use an unique id but this is just for didactic purposes :)
	filter := 
	// Query the DeleteOne() function with your filter and you are set!
	result, err := 
	if err != nil {
        return 
    }
	// return http 200 when all is fine with the result as the data payload
	return 
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/user/:name", getUser)
	app.Post("/user", createUser)
	app.Put("/user/:name", updateUser)
	// Add the Delete uri path with the "/user/:name"
	app.Delete("/user/:name", deleteUser)
}

// main 
func main (){
	// Fiber 
	app := fiber.New()
	setupRoutes(app)
	// Listen to port 3000
    app.Listen(":3000")
}
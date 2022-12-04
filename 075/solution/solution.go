package main

// In this exercise we will Stablish a connection with mongoDB - Atlas
// MongoDB is a non relational kind of db, and it's saas has a free tier. Find out more here.
// https://www.mongodb.com/docs/atlas/tutorial/deploy-free-tier-cluster/
// You can also use a docker container locally to connect, use whatever option suits you best.

// 1- You need to install these two libraries through the following commands, one is for the mongo driver and the other for a dotenv file
// go get go.mongodb.org/mongo-driver/mongo
// go get github.com/joho/godotenv
// 2- Once installed, create a file named ".env" containing your 
import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main (){

}
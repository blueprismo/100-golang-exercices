# Exercise 90: MongoDB Connection Setup

## Introduction to MongoDB

MongoDB is a popular NoSQL document database that stores data in flexible, JSON-like documents. It's designed for scalability, performance, and ease of development with modern applications.

## MongoDB Atlas

MongoDB Atlas is MongoDB's cloud database service offering:
- **Free Tier**: Perfect for learning and development
- **Managed Service**: No infrastructure management needed
- **Global Clusters**: Deploy worldwide
- **Built-in Security**: Authentication and encryption

## Go MongoDB Driver

The official MongoDB Go driver provides:
- Native Go API for MongoDB operations
- Connection pooling and management
- BSON encoding/decoding
- Comprehensive error handling

```go
import "go.mongodb.org/mongo-driver/mongo"
import "go.mongodb.org/mongo-driver/mongo/options"
```

## Environment Variables

Using `.env` files for configuration:
- Keeps sensitive data out of source code
- Different environments (dev, staging, prod)
- Easy configuration management

```go
import "github.com/joho/godotenv"

err := godotenv.Load()
uri := os.Getenv("MONGODB_URI")
```

## Connection Lifecycle

1. **Load Environment**: Read connection string from .env
2. **Create Client**: Initialize MongoDB client with URI
3. **Connect**: Establish connection to database
4. **Use Database**: Perform operations
5. **Disconnect**: Clean up resources

## Your Task

Look at the main.go file and complete the exercise by:
1. Installing required packages (mongo-driver and godotenv)
2. Loading environment variables from .env file
3. Establishing connection to MongoDB Atlas
4. Implementing proper connection cleanup

This exercise establishes the foundation for database operations in Go applications.
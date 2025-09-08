# Exercise 96: REST API with Fiber - Complete CRUD

## Fiber Web Framework

Fiber is an Express.js-inspired web framework for Go:
- **Fast Performance**: Built on Fasthttp
- **Express-like API**: Familiar syntax for web developers
- **Middleware Support**: Extensive middleware ecosystem
- **JSON Handling**: Built-in JSON parsing and response

## CRUD Operations

Complete CRUD (Create, Read, Update, Delete) operations:
- **CREATE**: POST /user - Add new user
- **READ**: GET /user/:name - Retrieve user
- **UPDATE**: PUT /user/:name - Modify user
- **DELETE**: DELETE /user/:name - Remove user

## HTTP Methods and Routes

```go
app.Get("/user/:name", getUser)       // Read
app.Post("/user", createUser)         // Create  
app.Put("/user/:name", updateUser)    // Update
app.Delete("/user/:name", deleteUser) // Delete
```

## Request/Response Patterns

- **Route Parameters**: `c.Params("name")` extracts URL parameters
- **Request Body**: `c.BodyParser()` parses JSON payloads
- **JSON Response**: `c.JSON()` sends JSON responses
- **Status Codes**: Proper HTTP status codes for each operation

## MongoDB Integration

Combining Fiber with MongoDB:
- Database operations in route handlers
- BSON for database queries and updates
- Error handling for database operations
- Structured JSON responses

## API Response Structure

Consistent response format:
```go
type UserResponse struct {
    Status  int        `json:"status"`
    Message string     `json:"message"`
    Data    *fiber.Map `json:"data"`
}
```

## Your Task

Look at the main.go file and complete the exercise by:
1. Implementing the deleteUser function
2. Creating proper BSON filters for deletion
3. Using MongoDB's DeleteOne operation
4. Returning appropriate HTTP responses

This exercise completes a full REST API with all CRUD operations using Fiber and MongoDB.
# Exercise 99: API Data Creation with POST

## HTTP POST Method

POST requests create new resources:
- **Request Body**: Data sent in HTTP body (usually JSON)
- **Content-Type**: `application/json` for JSON data
- **Status Codes**: 201 Created for successful creation
- **Response**: Return created resource with ID

## Request Body Parsing

Fiber's `BodyParser()` method:
- Automatically parses JSON to Go structs
- Handles content-type detection
- Validates JSON structure
- Returns parsing errors

```go
var user User
err := c.BodyParser(&user)
```

## JSON Struct Tags

Proper JSON mapping with struct tags:
```go
type User struct {
    Age  int    `json:"age"`
    Name string `json:"name"`
}
```

## Database Insertion Flow

1. **Parse Request**: Extract JSON from request body
2. **Validate Data**: Check required fields and formats
3. **Insert Document**: Add to database collection
4. **Return Response**: Confirm creation with inserted ID

## Error Scenarios

Handle common POST errors:
- **400 Bad Request**: Invalid JSON or missing fields
- **409 Conflict**: Duplicate resource
- **500 Internal Server Error**: Database errors
- **503 Service Unavailable**: System issues

## Testing POST Endpoints

Tools for testing POST APIs:
- **Postman**: GUI tool for API testing
- **curl**: Command-line HTTP client
- **HTTPie**: User-friendly HTTP client
- **Automated Tests**: Go testing with HTTP requests

## Your Task

Look at the main.go file and complete the exercise by:
1. Connecting to database collection
2. Parsing JSON request body into User struct
3. Creating new user document with parsed data
4. Inserting document and returning success response

This exercise completes the CREATE operation in a REST API, handling JSON input and database insertion.
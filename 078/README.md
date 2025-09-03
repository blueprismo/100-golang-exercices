# Exercise 78: Creating a JSON API

## Building REST APIs

APIs (Application Programming Interfaces) allow different systems to communicate. JSON APIs are the most common format for modern web services.

## JSON API Basics

A JSON API typically:
1. **Accepts** JSON requests
2. **Processes** data  
3. **Returns** JSON responses
4. **Uses** HTTP methods (GET, POST, PUT, DELETE)

## Setting Response Headers

For JSON responses, set the content type:
```go
w.Header().Set("Content-Type", "application/json")
```

## Returning JSON Data

```go
type Response struct {
    Message string `json:"message"`
    Data    Animal `json:"data"`
}

response := Response{Message: "success", Data: animal}
json.NewEncoder(w).Encode(response)
```

## HTTP Status Codes

- **200**: OK
- **201**: Created  
- **400**: Bad Request
- **404**: Not Found
- **500**: Internal Server Error

## Your Task

Create a JSON API with an Animal struct:
1. Define an Animal struct with Name and Type fields
2. Create an endpoint that returns JSON data
3. Set proper content type headers
4. Return structured JSON responses

This is the foundation for building REST APIs that other applications can consume.
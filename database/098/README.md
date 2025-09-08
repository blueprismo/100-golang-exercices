# Exercise 98: API Data Retrieval with Parameters

## URL Parameters in Fiber

Route parameters allow dynamic URLs where parts of the path are variables:
- **Definition**: `/user/:name` defines a parameter named "name"
- **Access**: `c.Params("name")` retrieves parameter value
- **Flexible**: Multiple parameters per route supported

```go
app.Get("/user/:id/posts/:postid", handler)
// Access: c.Params("id"), c.Params("postid")
```

## Database Queries with Parameters

Using route parameters in database queries:
1. Extract parameter from URL
2. Create database filter using parameter
3. Execute query with filter
4. Return results as JSON

## Error Handling Patterns

Proper API error handling:
- **404 Not Found**: Resource doesn't exist
- **500 Internal Server Error**: Database/system errors
- **400 Bad Request**: Invalid parameters
- **Consistent Format**: Standard error response structure

## JSON Response Structure

Structured API responses provide consistency:
```go
type UserResponse struct {
    Status  int        `json:"status"`
    Message string     `json:"message"`  
    Data    *fiber.Map `json:"data"`
}
```

## Case Sensitivity in Databases

MongoDB queries are case-sensitive by default:
- "John" ≠ "john" 
- Use exact matching for string fields
- Consider case-insensitive queries for search features
- Consistent data input validation important

## Your Task

Look at the main.go file and complete the exercise by:
1. Extracting name parameter from URL
2. Creating BSON filter for database query
3. Executing FindOne operation and decoding result
4. Implementing proper error handling and responses

This exercise demonstrates parameterized API routes and database integration with proper error handling.
# Exercise 100: API Data Updates with PUT

## HTTP PUT Method

PUT requests update existing resources:
- **Complete Updates**: Replaces entire resource
- **Idempotent**: Same request multiple times has same effect
- **Request Body**: New data in JSON format
- **URL Parameters**: Identify resource to update

## MongoDB Update Operations

MongoDB provides several update methods:
- **UpdateOne()**: Update single matching document
- **UpdateMany()**: Update multiple matching documents
- **ReplaceOne()**: Replace entire document
- **FindOneAndUpdate()**: Update and return document

## Update Documents with $set

The `$set` operator updates specific fields:
```go
update := bson.M{"name": user.Name, "age": user.Age}
result, err := collection.UpdateOne(ctx, filter, bson.M{"$set": update})
```

## Update vs Replace

- **Update ($set)**: Modifies specific fields, preserves others
- **Replace**: Completely replaces document content
- **Partial Updates**: Use PATCH method for partial updates
- **Full Updates**: Use PUT method for complete replacement

## Capped Collection Limitations

Capped collections have update restrictions:
- Documents cannot grow larger than original size
- Some update operations may fail
- Consider design implications for capped collections

## Error Handling for Updates

Common update scenarios:
- **404 Not Found**: Document doesn't exist
- **400 Bad Request**: Invalid update data
- **409 Conflict**: Update constraint violations
- **500 Internal Error**: Database operation failures

## Your Task

Look at the main.go file and complete the exercise by:
1. Creating BSON update document with new values
2. Setting up filter to identify document to update
3. Executing UpdateOne operation with $set operator
4. Returning appropriate success response

This exercise completes the UPDATE operation in a REST API, demonstrating document modification in MongoDB.

## Congratulations! 🎉

You've completed all 100 Go exercises! This final exercise wraps up your journey through:
- Basic Go concepts and syntax
- Concurrency with goroutines and channels  
- Time handling and formatting
- JSON processing and APIs
- Context management
- Interfaces and type systems
- HTTP servers and web frameworks
- Testing strategies and tools
- Database operations with MongoDB

You now have a solid foundation in Go programming and are ready to build real-world applications!
# Exercise 92: BSON and Document Insertion

## What is BSON?

BSON (Binary JSON) is MongoDB's binary serialization format for storing documents. It extends JSON with additional data types like ObjectId, Date, and Binary data while maintaining JSON-like simplicity.

## Go Structs and BSON

Go structs map naturally to BSON documents:
- Struct fields become document fields
- Field tags control serialization
- Embedded structs create nested documents

```go
type User struct {
    Name string `bson:"name"`
    Age  int    `bson:"age"`
}
```

## Document Insertion

MongoDB provides several insertion methods:
- **InsertOne()**: Insert single document
- **InsertMany()**: Insert multiple documents
- **ReplaceOne()**: Replace entire document

```go
result, err := collection.InsertOne(context.TODO(), user)
insertedID := result.InsertedID
```

## BSON Tags

Common BSON struct tags:
- `bson:"fieldname"`: Map to specific field name
- `bson:",omitempty"`: Omit empty values
- `bson:"-"`: Exclude field from serialization
- `bson:",inline"`: Inline embedded struct fields

## Context Usage

MongoDB operations require context for:
- **Cancellation**: Cancel long-running operations
- **Timeouts**: Set operation deadlines
- **Metadata**: Pass request-scoped values

## Your Task

Look at the main.go file and complete the exercise by:
1. Defining User struct with proper BSON tags
2. Creating user instances with data
3. Inserting documents into MongoDB collection
4. Understanding BSON serialization in Go

This exercise teaches you document structure and basic insertion operations in MongoDB.
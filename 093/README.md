# Exercise 93: Document Queries and Unmarshaling

## Querying Documents

MongoDB uses filters to query documents. In Go, filters are typically created using `bson.D` (ordered) or `bson.M` (unordered) types:

```go
// Find by exact match
filter := bson.D{{"name", "John"}}

// Find by multiple criteria
filter := bson.D{{"name", "John"}, {"age", 25}}
```

## BSON Types in Go

The MongoDB Go driver provides several BSON types:
- **bson.D**: Ordered document (slice of key-value pairs)
- **bson.M**: Unordered document (map)
- **bson.A**: Array (slice)
- **bson.E**: Element (key-value pair)

## FindOne Operation

`FindOne()` retrieves a single document matching the filter:
- Returns first matching document
- Requires `Decode()` to unmarshal into Go struct
- Returns error if no document found

```go
var user User
err := collection.FindOne(context.TODO(), filter).Decode(&user)
```

## Unmarshaling Process

1. **Query**: MongoDB returns BSON document
2. **Decode**: Convert BSON to Go struct
3. **Type Mapping**: BSON types map to Go types
4. **Field Matching**: BSON field names match struct tags

## Error Handling

Common query errors:
- **mongo.ErrNoDocuments**: No document matches filter
- **Decode errors**: BSON to struct conversion issues
- **Connection errors**: Database connectivity problems

## Your Task

Look at the main.go file and complete the exercise by:
1. Creating query filters using bson.D
2. Executing FindOne operations
3. Decoding BSON results into Go structs
4. Handling query results and errors

This exercise demonstrates document retrieval and BSON unmarshaling in MongoDB operations.
# Exercise 94: Bulk Insert Operations

## Bulk Operations

Bulk operations allow inserting multiple documents in a single database round-trip, providing better performance than individual insertions.

## InsertMany()

`InsertMany()` accepts a slice of documents:
- More efficient than multiple `InsertOne()` calls
- Atomic operation (all succeed or all fail)
- Returns slice of inserted IDs
- Supports ordered and unordered insertions

```go
docs := []interface{}{
    User{Name: "Alice", Age: 30},
    User{Name: "Bob", Age: 25},
}
result, err := collection.InsertMany(context.TODO(), docs)
```

## Interface{} Usage

MongoDB operations often use `interface{}` for flexibility:
- Accepts any document type
- Allows mixed document structures
- Requires type assertion for specific operations
- Common pattern in database operations

## Bulk Operation Benefits

- **Performance**: Fewer network round-trips
- **Atomicity**: All operations succeed or fail together
- **Efficiency**: Reduced connection overhead
- **Throughput**: Higher data ingestion rates

## Ordered vs Unordered

- **Ordered**: Stops on first error, maintains insertion order
- **Unordered**: Continues on errors, may reorder for performance

```go
opts := options.InsertMany().SetOrdered(false)
result, err := collection.InsertMany(ctx, docs, opts)
```

## Use Cases

- **Data Migration**: Moving large datasets
- **Batch Processing**: Processing multiple records
- **Import Operations**: Loading external data
- **Seeding**: Populating test or initial data

## Your Task

Look at the main.go file and complete the exercise by:
1. Creating a slice of user documents with interface{} type
2. Including a user named "Rose" as specified
3. Using InsertMany() for bulk insertion
4. Understanding bulk operation results

This exercise teaches efficient bulk data insertion patterns in MongoDB.
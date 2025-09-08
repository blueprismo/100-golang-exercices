# Exercise 91: MongoDB Collections

## Understanding MongoDB Collections

Collections in MongoDB are equivalent to tables in relational databases. They group related documents together and can be created explicitly or automatically when first accessed.

## Collection Operations

MongoDB provides several collection management operations:
- **CreateCollection()**: Explicitly create collections with options
- **Collection()**: Get reference to existing or new collection
- **ListCollections()**: View all collections in database
- **Drop()**: Remove collections

## Capped Collections

Capped collections have special characteristics:
- **Fixed Size**: Maximum storage size defined at creation
- **FIFO Behavior**: Oldest documents removed when size limit reached
- **High Performance**: Optimized for high-throughput operations
- **No Updates**: Documents cannot be updated to larger sizes

```go
opts := options.CreateCollection()
opts.SetCapped(true)
opts.SetSizeInBytes(1048576) // 1MB
```

## Collection Options

Common collection creation options:
- **Capped**: Create fixed-size collection
- **Size**: Maximum size in bytes
- **Max**: Maximum number of documents
- **Validator**: Document validation rules

## Use Cases for Capped Collections

- **Logging**: High-volume log data with automatic rotation
- **Metrics**: Time-series data with size constraints
- **Cache**: Fixed-size cache with automatic eviction
- **Event Streams**: Real-time event processing

## Your Task

Look at the main.go file and complete the exercise by:
1. Setting up collection creation options
2. Creating a capped collection with size limits
3. Getting collection references for operations
4. Understanding collection configuration

This exercise introduces you to MongoDB collection management and specialized collection types.
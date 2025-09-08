# Exercise 95: MongoDB Aggregation Pipeline

## What is Aggregation?

Aggregation in MongoDB processes data records and returns computed results. It's similar to SQL's GROUP BY operations but more powerful and flexible.

## Aggregation Pipeline

The aggregation pipeline consists of stages that process documents:
- Documents pass through stages sequentially
- Each stage transforms the document stream
- Stages can filter, group, sort, and compute values

```go
pipeline := mongo.Pipeline{
    {{"$match", bson.D{{"status", "A"}}}},
    {{"$group", bson.D{{"_id", "$category"}, {"total", bson.D{{"$sum", "$amount"}}}}}},
}
```

## Common Aggregation Stages

- **$match**: Filter documents (like WHERE clause)
- **$group**: Group documents and perform calculations
- **$sort**: Sort documents
- **$project**: Select/transform fields
- **$limit**: Limit number of results
- **$skip**: Skip documents

## Group Stage Operations

The `$group` stage performs calculations:
- **$sum**: Calculate totals
- **$avg**: Calculate averages  
- **$min/$max**: Find minimum/maximum values
- **$count**: Count documents
- **$push**: Create arrays of values

## BSON Document Structure

Group stage example:
```go
groupStage := bson.D{
    {"$group", bson.D{
        {"_id", "$name"},                              // Group by name
        {"average_age", bson.D{{"$avg", "$age"}}},    // Calculate average age
        {"count", bson.D{{"$sum", 1}}},               // Count occurrences
    }},
}
```

## Your Task

Look at the main.go file and complete the exercise by:
1. Creating a group stage with proper BSON structure
2. Using aggregation operations like $avg and $sum
3. Running aggregation pipeline with Aggregate()
4. Processing aggregation results

This exercise introduces MongoDB's powerful aggregation framework for data analysis and reporting.
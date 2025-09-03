# Exercise 82: Gin Query Parameters

## Query String Parameters

Query parameters are key-value pairs in the URL after the `?` symbol. They're used for filtering, pagination, and optional parameters.

## Query Parameter Examples

```
http://example.com/albums?artist=Beatles&year=1969
http://example.com/search?q=go+programming&limit=10
```

## Accessing Query Parameters in Gin

```go
r.GET("/albums", func(c *gin.Context) {
    artist := c.Query("artist")        // Required parameter
    year := c.DefaultQuery("year", "") // Optional with default
    
    if artist == "" {
        c.JSON(400, gin.H{"error": "artist parameter required"})
        return
    }
    
    // Use parameters to filter results
})
```

## Query Parameter Methods

- **c.Query("key")**: Get parameter value (empty string if not found)
- **c.DefaultQuery("key", "default")**: Get parameter with default value
- **c.GetQuery("key")**: Get parameter and boolean indicating if present

## Common Use Cases

- **Filtering**: `?category=electronics&price=100`
- **Pagination**: `?page=2&limit=20`
- **Search**: `?q=search+term`
- **Sorting**: `?sort=name&order=asc`

## Your Task

Create an endpoint that handles query parameters:
1. Accept query parameters (like artist name, year, etc.)
2. Use parameters to filter or search albums
3. Provide default values for optional parameters
4. Return filtered results based on query parameters
5. Handle missing required parameters gracefully

This demonstrates building flexible APIs that accept optional filtering criteria.
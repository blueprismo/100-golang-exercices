# Exercise 80: Gin POST Requests

## Handling POST Requests

POST requests are used to send data to the server (creating new resources, form submissions, etc.). Gin makes handling POST requests and JSON data very easy.

## POST Route in Gin

```go
r.POST("/albums", func(c *gin.Context) {
    // Handle POST request
})
```

## Binding JSON Data

Gin can automatically bind JSON request bodies to Go structs:

```go
type Album struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Artist string `json:"artist"`
}

func createAlbum(c *gin.Context) {
    var newAlbum Album
    
    if err := c.ShouldBindJSON(&newAlbum); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    
    // Process the album data
    c.JSON(201, newAlbum)
}
```

## HTTP Methods in Gin

- **GET**: `r.GET("/path", handler)`
- **POST**: `r.POST("/path", handler)`
- **PUT**: `r.PUT("/path", handler)`
- **DELETE**: `r.DELETE("/path", handler)`

## Testing POST Requests

Use curl to test:
```bash
curl -X POST http://localhost:8080/albums \
  -H "Content-Type: application/json" \
  -d '{"id":"1","title":"Blue Album","artist":"Weezer"}'
```

## Your Task

Create a Gin server that:
1. Defines an album struct with JSON tags
2. Handles POST requests to create new albums  
3. Binds JSON request data to the struct
4. Returns appropriate status codes
5. Handles binding errors gracefully

This demonstrates building REST APIs that accept data.
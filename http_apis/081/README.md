# Exercise 81: Gin Path Parameters

## URL Path Parameters

Path parameters allow you to capture parts of the URL as variables. This is essential for REST APIs where you need to identify specific resources.

## Path Parameter Syntax

```go
// Route with parameter
r.GET("/albums/:id", getAlbumByID)

// In handler function
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")  // Get the :id parameter
    // Use id to find specific album
}
```

## REST API Patterns

Common RESTful URL patterns:
- **GET /albums**: Get all albums
- **GET /albums/:id**: Get specific album
- **POST /albums**: Create new album
- **PUT /albums/:id**: Update specific album
- **DELETE /albums/:id**: Delete specific album

## Finding Resources

Typically you'd search a database or slice:
```go
func findAlbumByID(id string) (*Album, bool) {
    for _, album := range albums {
        if album.ID == id {
            return &album, true
        }
    }
    return nil, false
}
```

## HTTP Status Codes

- **200**: Found and returning resource
- **404**: Resource not found

## Your Task

Create a GET endpoint that:
1. Accepts an ID as a path parameter (`/albums/:id`)
2. Searches for an album with that ID
3. Returns the album if found (200 status)
4. Returns 404 error if not found
5. Uses proper REST conventions

This demonstrates building resource-specific API endpoints.
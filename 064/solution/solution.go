// Exercise: Create an API with GIN framework
package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

// album represents data about a record album.
type album struct {
  ID     string  `json:"id"`
  Title  string  `json:"title"`
  Artist string  `json:"artist"`
  Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
  {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
  {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
  {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// Then, let's create a handler function ((a normal function, which then we will wrap :) ))
// name it getAlbums, and it's only argument named 'c' will be a pointer type of gin.Context.
func getAlbums(c *gin.Context) {
  // use the IndentedJSON function with the gin context, and pass it an http status ok, and the albums array
  // more info: https://pkg.go.dev/github.com/gin-gonic/gin#Context.IndentedJSON
  c.IndentedJSON(http.StatusOK, albums)
}

func main() {  
  // we will registrer a handler (or router) with gin.Default()
  router := gin.Default()
  
  // here access the router .GET http verb for the request. 
  // https://pkg.go.dev/github.com/gin-gonic/gin#readme-using-get-post-put-patch-delete-and-options
  // the first argument will be the URI (or pattern) and the second one the getAlbums handler function we defined before
  router.GET("/albums", getAlbums)

  // run the server with the Run() function
  router.Run("localhost:8080")
  
}
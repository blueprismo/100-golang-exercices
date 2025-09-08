// Exercise: Create an API with GIN framework - GET request to specific ITEM
package main

// In this exercise what we want to do is get an specific album by it's ID
// The path we will want to go is "/album/$ID" and it will return our album with the ID (if it exists!!!)

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

func getAlbums(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums( c *gin.Context){
  var newAlbum album
  if err := c.BindJSON(&newAlbum); err != nil {
    return
  }
  albums = append(albums,newAlbum)
  c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getSpecificAlbum(c *gin.Context){
  // get the ID in a variable

  // check if it matches any of the slice ID (you have to iterate through the array)

  // return some error message when album isn't found!

}

func main() {  
  // we will registrer a handler (or router) with gin.Default()
  router := gin.Default()
  router.GET("/albums", getAlbums)
  router.POST("/albums", postAlbums)
  // create the route specific to get the /album/$ID (https://github.com/gin-gonic/gin#parameters-in-path)
  
  // run the server with the Run() function
  router.Run("localhost:8080")
}
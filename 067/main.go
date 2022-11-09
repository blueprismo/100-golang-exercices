// Exercise: Create an API with GIN framework - Query string parameters
package main

// In this exercise what we want to do is handle query string parameters
// Query string parameters are the ones you may find in the url like this:
// http://example.com/welcome?name=enin&surname=tolstoy
//                            <---1--->&<---2---------> .  2 parameters in this case


// We will handle them with gin, and we will try to access an album the same way in the previous exercise, but with it's query string parameters
// We will modify the first function to handle both ways (with and without query string parameters)
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

// IN THIS EXERCISE YOU HAVE TO MODIFY THIS FUNCTION!!
func getAlbums(c *gin.Context) {
  // get the query string (look up for the .Query() function in gin :) )
  // our QueryString key will be "id" and the value, the ID we want to get
  
  // if the id is null, serve all the albums (like in the first gin API exercise)
  

  // else, loop over the array to serve the matching album.
    
  // return some error message when album isn't found inside this else block!!
  
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
  var id string = c.Param("id")
  for _, album := range albums {
		if id == album.ID {
			c.IndentedJSON(http.StatusOK, album)
      return
		}
	}
  c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}


func main() {  
  // we will registrer a handler (or router) with gin.Default()
  router := gin.Default()

  router.GET("/albums", getAlbums)
  router.POST("/albums", postAlbums)
  router.GET("/albums/:id", getSpecificAlbum)

  // run the server with the Run() function
  router.Run("localhost:8080")
}
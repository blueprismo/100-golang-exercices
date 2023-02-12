package main

// In this sets of exercises we are going to learn and use Heimdall
// To install Heimdall we need to run `go get -u github.com/gojek/heimdall/v7`
// Then import the package with import "github.com/gojek/heimdall/v7/httpclient" 

import (
  "time"
  "fmt"
  "io"
  
  ) 

// In this first exercise, we are going to make a simple http request to the google.com webpage
func main () {
  // Create a new HTTP client with a default timeout
  
  
  // Use the clients GET method to create and execute the request
  
  // Heimdall returns the standard *http.Response object
  body, err := io.ReadAll(res.Body)
  fmt.Println(string(body))
}
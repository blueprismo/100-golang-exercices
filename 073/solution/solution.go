package main

// In this sets of exercises we are going to learn and use Heimdall
// To install Heimdall we need to run `go get -u github.com/gojek/heimdall/v7`
// Then import the package with import "github.com/gojek/heimdall/v7/httpclient" 

import (
  "time"
  "fmt"
  "io"
  "github.com/gojek/heimdall/v7/httpclient"
  ) 

// In this first exercise, we are going to make a simple http request to the google.com webpage
func main () {
  // Create a new HTTP client with a default timeout
  timeout := 1000 * time.Millisecond
  client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

  // Use the clients GET method to create and execute the request
  res, err := client.Get("http://google.com", nil)
  if err != nil{
  	panic(err)
  }

  // Heimdall returns the standard *http.Response object
  body, err := io.ReadAll(res.Body)
  fmt.Println(string(body))
}
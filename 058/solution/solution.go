// Exercise: Set up a simple HTTP Server

// We will use the net/http library
// https://pkg.go.dev/net/http

// There are multiple ways of creating a http server in GO
// In this task, we are going to crete a http in your preferred way, so run free with the documentation :)

package main

import "net/http"
import "fmt"
import "html"
import "log"

func main() {

  http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  log.Fatal(http.ListenAndServe(":8080", nil))
}
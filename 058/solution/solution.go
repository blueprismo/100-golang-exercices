// Exercise: Set up a simple HTTP Server

// We will use the net/http library
// https://pkg.go.dev/net/http

// Like with ExpressJS, we will make a web server, and the web server will serve the "/bar" route
// and the response should be "Hello, /var". BUT don't hardcode the URI!
// Changing the http.HandleFunc 1st variable (the URI), the message served should also change.
// Example: If I set my webserver at /newpath, my response will be "Hello, /newpath"

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
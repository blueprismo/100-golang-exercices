package main

// In this exercise we will create a microservice that counts the number of requests.
// Tip: This exercise is ressemblant to the one in exercise 62, but better! (And improved!)
import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// For simplicity we will obviate the Mutex and the defer in the next function. We will get to that later
// Create a countHandler struct with a variable n type int.
type countHandler struct {
	mu sync.Mutex // guards n
	n  int
}

// Create the http handler function for creating a counter (tip: take a look at exercise 62!)
func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  log.Println("received the", h.n + 1," request")
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

// We will use the default serve mux here
func main() {
	http.Handle("/count", new(countHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
package main
// In this exercise we are going to talk about mutexes and wait groups. As GO is useful for concurrent programs, we have to be aware about concurrent executions problems and race conditions

import (
	"fmt"
	"sync"
	"time"
)
// Imagine we have a single counter. But we have 5 goroutines that use and increment that counter. If we don't put locks between the counters
// We assume the risk of some of the counters acting with false information, and therefore making the counter useless.
// Let's make use of mutexes
// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	n  int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc() {
	// Lock so only one goroutine at a time can access the map c.v.
	c.mu.Lock()
	// Here is where we have a critical section! https://en.wikipedia.org/wiki/Critical_section
	c.n++
	//Unlock
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value() int {
	// Lock so only one goroutine at a time can access the map c.v.
	c.mu.Lock()
	//Defer the mutex Unlock, as the deferred functions will be executed after the return is evaluated :)
	defer c.mu.Unlock()
	// return the total counters
	return c.n
}

func main() {
	// Create a SafeCounter variable named "c"
	var c SafeCounter 
	// Create a waitgroup named "wg"
	var wg sync.WaitGroup
	// Add 5 [CONCURRENT] goroutines to the waitgroup with the Add() function
	wg.Add(5)

	// Create a loop for the 5 gogroutines
	for i := 0; i < 5; i++ {
		//Create the go routine here
		go func() {
			// Defer the waitgroup.Done() call
			defer wg.Done()
			// Create a loop that iterates from 0 to 100.000
			for j := 0; j < 100000; j++ {
				// Use the Inc() function to increment the value of the counter!
				c.Inc()
			}
		}()
	}
	// wait for the goroutines to end
	wg.Wait()

	time.Sleep(time.Second)
	// Print the value... notice the value returned by the counter!!!!!! Is it 100.000 or 500.000?
	fmt.Println(c.Value())
}
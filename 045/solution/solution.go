// Exercise: Tickers

// Create a goroutine with a infinite loop
// The condition for that infinite loop will be a range that goes over the "time.Tick(time.Second * 1)"
// And at every Tick, we will print "Tick"
// In the main function, call the go routine :)

package main

import "fmt"
import "time"

func task() {
    for range time.Tick(time.Second *1){
        fmt.Println("Tick ")
    }
}

func main() {
    go task()
    time.Sleep(time.Second * 5)
}

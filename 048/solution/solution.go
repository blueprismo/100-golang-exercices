// Exercise: Time

// We will have 2 time variables
// You will have to: 
// 1- Get the difference with the Sub() function
// 2- See if they are equal with the Equal() funciton
// 3- See what comes after the other with the After() function

package main

import (
    "fmt"
    "time"
)

func main() {
  start := time.Date(2020, 2, 1, 3, 0, 0, 0, time.UTC)
  end   := time.Date(2021, 2, 1, 12, 0, 0, 0, time.UTC)

  difference := end.Sub(start) 
  fmt.Printf("The difference is %v: ", difference)

  equal := start.Equal(end) 
  if equal {
    fmt.Printf("They are equal")
  } else {
    fmt.Printf("They are NOT equal")
  }
  
  whatAfter := start.After(end) 
  fmt.Printf("It's %t that the start date comes after the end date", whatAfter)
}

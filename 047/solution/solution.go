// Exercise: Time

// Add 20 minutes to the current time (shown in UTC)

package main

import (
    "fmt"
    "time"
)

func main() {
  current := time.Now().UTC()
  inTenMinutes := current.Add(time.Minute * 20 )
  
  fmt.Println(inTenMinutes)
}

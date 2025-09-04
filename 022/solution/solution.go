// Exercise: Random
// Create a program that mimics a dice roll (a 6-face dice)
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func random(min int, max int) int {
    return rand.Intn(max-min) + min
}

func main() {
    rand.Seed(time.Now().UnixNano())
    randomNum := random(1, 7)
    fmt.Printf("Rolled dice: %d\n", randomNum)
}
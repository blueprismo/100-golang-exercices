package main

// import the "testing" library
import (
    "testing"
    "math/rand"
)

// Here, create a function named BenchmarkXxxxxxx with the signature (b *testing.B)
func BenchmarkRandInt (b *testing.B){
    // make a for loop from 0 to b.N 
    for i := 0; i < b.N ; i++{
        rand.Int()
    }
} 
// Now run go test -bench . /path/to/this/file
// and see what happens!
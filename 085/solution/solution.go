// TESTING!
package main

// In this exercise, we are going to create 3 tests for the reverse function, that reverses a string
// if it gets "Hello" it will return "olleH"

import "fmt"

func Reverse(s string) string {
  b := []byte(s)
  for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
      b[i], b[j] = b[j], b[i]
  }
  return string(b)
}

func main() {
  input := "The quick brown fox jumped over the lazy dog"
  rev := Reverse(input)
  doubleRev := Reverse(rev)
  fmt.Printf("original: %q\n", input)
  fmt.Printf("reversed: %q\n", rev)
  fmt.Printf("reversed again: %q\n", doubleRev)
}
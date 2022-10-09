// Exercise: Check if a file exists
// Tip: use the "os" package

package main

import "fmt"
import "os"

func main () {
	// Here goes your code
	if _, err := os.Stat("file-exists.go"); err == nil {
		fmt.Printf("File exists\n");  
	  } else {
		fmt.Printf("File does not exist\n");  
	  }
}
// Exercise: Read a file
// Tip: use the "io and os" package

// beware: You should run this code where the read file is, or reference it!
package main

import "fmt"
import "io"
import "os"

func main () {
  // Open the file
	file, err := os.Open()
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the entire file content using io.ReadAll
	
	// Print the content

}

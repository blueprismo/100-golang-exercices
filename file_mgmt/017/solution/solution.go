// Exercise: Read a file
// Tip: use the "io/ioutil" package

// beware: You should run this code where the read file is, or reference it!
package main

import "fmt"
import "io"
import "os"

func main () {
  // Open the file
	file, err := os.Open("read.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the entire file content using io.ReadAll
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the content
	fmt.Println(string(content))

}

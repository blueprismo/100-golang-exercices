// Exercise: Read a file

// beware: You should run this code where the read file is, or reference it!
package main

import (
	"log"
  	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}

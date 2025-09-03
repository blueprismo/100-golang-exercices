// Exercise: Read a file
// Tip: use the "io/ioutil" package

// beware: You should run this code where the read file is, or reference it!
package main

import (
	"log"
  	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("/tmp/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data)) // or os.Stdout.Write(data)
}
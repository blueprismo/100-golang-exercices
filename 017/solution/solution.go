// Exercise: Read a file
// Tip: use the "io/ioutil" package

// beware: You should run this code where the read file is, or reference it!
package main

import "fmt"
import "io/ioutil"

func main () {
	// Here goes your code
	b, err := ioutil.ReadFile("read.go")
    // can file be opened?
    if err != nil {
      fmt.Print(err)
    }

    // convert bytes to string
    str := string(b)

    // show file data
    fmt.Println(str)
}
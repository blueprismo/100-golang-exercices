# Exercise: Arrays II

A string variable is an array of tinier elements of type `char`.

For example the word `TREE` is an array of 4 characters, each character will have an assigned index.
The first element (with index 0) is the letter 'T', the second one (with index 1) is the letter 'R' and so on.

```txt
T - R - E - E
0 - 1 - 2 - 3
```

Exercise: With a single string variable named helloWorld, print only the first character.

- string[n] is how you should access the value, with n being the index number.

```go
package main

import "fmt"

func main () {
  var helloWorld string
  helloWorld = "Hello World!"
  // Here goes your code
  fmt.Println()
}
```

<details>
<summary> Solution: </summary>

```go
package main

import "fmt"

func main() {
  // Creating new variable called helloWorld
  var helloWorld string
  helloWorld = "Hello World!"
  // Print the first letter
  fmt.Println(helloWorld[0])
}

// To run the program:
// - go run solution.go
```

</details>

# Exercise: With a single string variable, use access the first character with the string index

- This is only valid with ASCII characters, and the print value will be the ASCII number
- string[n] is how you should access the value

```golang
package main

import "fmt"

func main () {
	// Here goes your code
	var string1 string
	// ...
}
```

<details>
<summary> Solution: </summary>

```golang
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

# Exercise: Concatenate two string variables and print it's result

- One variable should be called "helloWorld" and the other "itsmemario"
- There should be a space between them

```golang
package main

func main() {
	// Here goes your code
	var string1, string2 string
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
	var helloWorld, itsMeMario string
	helloWorld = "Hello World!"
	itsMeMario = "It's a me, Mario"
	// Print the variable
	fmt.Println(helloWorld + " " + itsMeMario)
}

// To run the program:
// - go run solution.go
```

</details>

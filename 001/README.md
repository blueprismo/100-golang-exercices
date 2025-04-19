# Exercise: Run a program that prints "Hello World!"

- Use the "fmt" Package
- Print "Hello World!" to your standard output
- Run the program
- Make it an executable

```golang
package main

import "fmt"

func main () {
  // Here goes your code
}
```

<details>
<summary> Solution: </summary>

```golang
package main

import "fmt"

func main () {
	fmt.Println("Hello World!")
}

// To run the program:
// - go run solution.go

// To build an executable file
// - go build solution.go
```

</details>

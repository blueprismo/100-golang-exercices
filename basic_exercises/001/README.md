# Exercise: Run a program that prints "Hello World!"

Usually the first exercise while learning a programming language is printing ["Hello World!"](https://en.wikipedia.org/wiki/%22Hello,_World!%22_program) into the standard output.

You can use the [fmt.Println()](https://pkg.go.dev/fmt#Println) function to print a string into the standard output.

To run a go programm, you can spawn a terminal into your working directory and type `go run .`
To build an executable of that program, the `go build .` command will build a binary with your program that you can execute.

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

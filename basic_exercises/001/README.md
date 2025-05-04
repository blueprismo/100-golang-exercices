# Exercise: fmt.Println()

Usually the first exercise while learning a programming language is printing ["Hello World!"](https://en.wikipedia.org/wiki/%22Hello,_World!%22_program) into the standard output. Our goal here will be the same, but if you feel it's a hard start, there's more briefing for 'ya!

## Minimum amount of code

According to the [golang specification](https://go.dev/ref/spec#Program_execution)
> A complete program is created by linking a single, unimported package called the main package with all the packages it imports, transitively. The main package must have package name main and declare a function main that takes no arguments and returns no value.

For simplicity's sake, the minimum amount of code you will need to run your code is:

```golang
package main

func main () {

}
```

The code above executes nothing.  
`package main` just states the go compiler is a standalone executable.  
`func main () {}` is the entrypoint of the program, go will run this function when you execute the compiled binary.  

## Executing your code

To build an executable of a program, spawn a terminal into your working directory and type `go build .` command will build a binary with your program that you can execute.

Do we always need to build binaries after every execution?
There's no need for it as go provides the `go run .` command, which behind the scenes builds a temporary executable, it runs it, and then it deletes it.

## Exercise

In this exercise you will learn to import packages. You can import public libraries from other developers like you can import libraries that are in the go [standard library](https://pkg.go.dev/std)

To import a library, you just need to type `import <library_name>`, for example `import fmt` outside the main function.

*The fmt library implements formatted Input/Output functions, for more info: https://pkg.go.dev/fmt*

You can use the [fmt.Println()](https://pkg.go.dev/fmt#Println) function to print a string into the standard output.

Exercise: Run a program that prints "Hello World!"

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

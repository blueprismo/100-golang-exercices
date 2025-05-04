# Exercise: Reference and pointers

## Understanding Variable References and Indirection in Go

When you create a variable in Go, the computer doesn't just remember the name — it stores the value of that variable somewhere in the computer's memory. The name of the variable is just a convenient label we use in code, but behind the scenes, Go tracks where in memory the actual data lives.

This concept is called [indirection](https://en.wikipedia.org/wiki/Indirection)

## Memory Address and the & Operator

Every piece of data stored in memory has an address — think of it like a house address. The `&` operator in Go lets us find out where a variable is stored by giving us its memory address.

```go
package main

import "fmt"

func main() {
  var greeting string = "hello"

  fmt.Println("The value of greeting:", greeting)
  // Output: The value of greeting: hello

  fmt.Println("The address of greeting:", &greeting)
  // Output: The address of greeting: 0xc000010230 (your address will be different)
}
```

`greeting` stores the string "hello".
`&greeting` gives us the memory address where that string is stored.

Sometimes, you want to work with the memory address directly and this is where the concept of [pointers](https://en.wikipedia.org/wiki/Pointer_(computer_programming)) comes in. A pointer is just a variable that stores a memory address.

We’ll cover pointers in more detail later, but for now, remember:

`&x` means “give me the address of x”

This is the foundation of indirection, or accessing a value indirectly through its address

## Getting the user input

In order to get a user input we need:

1. A variable to store the input.
2. A way to tell Go where (in memory) to put the user’s input.

```go
var address string
fmt.Println("Please enter your address:")
fmt.Scanln(&address)
```

This code snippet tells Go: “Store the user's input into the variable `address` by writing to this memory location.”
You can think of `&name` like giving someone directions to your mailbox — it tells Go where to deliver the value.

You can read more about the Scanln function in: https://pkg.go.dev/fmt#hdr-Scanning

Exercise: Using only the fmt package, ask a user for it's name and then for it's surname

- Store it in 2 variables called "name" and "surname"
- After user has entered the data, print it to the standard output

```golang
package main

import "fmt"

func main () {
	var name, surname string
	// Here goes your code
	
}
```

<details>
<summary> Solution: </summary>

```golang
package main

import "fmt"

func main () {
	var name, surname string
	// Here goes your code
	fmt.Println("Please enter your name")
	fmt.Scanln(&name)
	fmt.Println("Please enter your surname")
	fmt.Scanln(&surname)

	fmt.Printf("Your name is: " + name + " " + surname)
}

```

</details>

# Exercise: Pointers II

A pointer is a variable that stores the memory address of another variable or object.
It's like a street address, pointing to the location where data is stored in the computer's memory. Pointers are useful for indirect memory access and memory manipulation.

The type `*T` is a pointer to a `T` value. Its zero value is `nil`.
`var p *int`

The `&` operator generates a pointer to its operand, in other words: To access the memory address of a variable you can prepend that variable with the `&` sign.

```go
// declare a pointer "p" of type int
var p *int
// declare i (type int) and assign the value 42
i := 42

// assign the address of i to the pointer named "p"
p = &i

// The * operator denotes the pointer's underlying value.
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p

// This is known as "dereferencing" or "indirecting".
```

- Declare a pointer variable of type int32
- Declare the address of the var "x"
- Save the address of the var "x" in the pointer variable

```golang
package main

import "fmt"

func main () {
    var x int32 = 5
	// Your code goes here

}
```

<details>
<summary> Solution: </summary>

```golang
package main

import "fmt"

func main () {
    var x int32 = 5
	// Your code goes here
	var pointerX *int32 = &x

	fmt.Println("Value of x: %d", x)
	fmt.Println("Memory address of x: %d", &x)
	fmt.Println("Pointer value: %d", pointerX)
}
```

</details>

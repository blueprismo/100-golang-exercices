# Exercise: Variables

## Variables in golang

A variable is an object that stores a value.

For example, the variable `name` can store the "Elvis" name.
The variable `number` can store 1,2,3

Variables in golang are statically-typed, meaning that once a variable type is defined, it can only store data of that type.
Exercise: Create a program that puts the "Hello World!" string inside a string variable called "helloWorld"

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
    /* 
    This is a 
    multiline
    comment
    */
	// Printing my name
	fmt.Println("My Name is John")
	// Printing my address!
	fmt.Println("My address is: Summs Rift 42")
}
```

</details>

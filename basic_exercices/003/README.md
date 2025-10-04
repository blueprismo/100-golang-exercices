# Exercise: Variables

## Variables in golang

A variable is an object that stores a value.

For example, the variable `name` can store the "Elvis" name.
The variable `number` can store 1, 20 or 500.
The variable `isHuman` can store `true` or `false`.

Variables in golang are statically-typed, meaning that once a variable type is defined, it can only store data of that type.

The variable `name` of type `string` contains a string.
The variable `number` of type `int` will contain an integer number.
The variable `isHuman` of type `bool` will contain a boolean.
And so on...

Variables are stored in the memory of the computer, but we will talk later about how they are stored or referenced.

## The shorthand :=

The `:=` syntax is shorthand for declaring and initializing a variable, e.g. for `var f string = "apple"` is equivalent to `f := "apple"`.
This syntax is only available inside functions, you can read more about it [here](https://go.dev/ref/spec#Short_variable_declarations).

Exercise: Create a program that puts the `Hello World!` string inside a string variable called `helloWorld`

Tip: String values are always wrapped in double-quotes `"hello, I am a string!"`

```go
package main

import "fmt"

func main () {
  // Here goes your code
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
  // Print the variable
  fmt.Println(helloWorld)
}

// To run the program:
// - go run solution.go
```

</details>

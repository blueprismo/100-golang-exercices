# Exercise: Comments

Comments in code are for explanatory purposes, if some other programmer is working with our code we can give a briefly explanation about what our code does. The comments are ignored by the compiler, it's only useful for the other humans to better understand the code.

## There are two types of comments in golang

Inline comments:
They only comment a single line, in order to comment a line in go, the double slash `//` is added at the beginning of the line.

Multi-line comments or block comments:
They comment multiple lines, they use the `/* <COMMENT> */` syntax.

- Use the fmt package

Exercise: Create a program that shows your name, address and also has some comments in it

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

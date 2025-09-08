# Exercise: Functions III

Variadic functions can be called with any number of arguments.
In golang, there is a term called Ellipsis which are three dots `...`, which is used for creating variadic functions.

According to the [golang spec](https://go.dev/ref/spec#Function_types):
> The final incoming parameter in a function signature may have a type prefixed with .... A function with such a parameter is called variadic and may be invoked with zero or more arguments for that parameter.

A function that is declared with an Ellipsis can take any nymber of arguments

```go
package main
  
import "fmt"
  
func main() {
  // 0 arguments 
    sayHello()
  // 1 argument
    sayHello("Rahul")
  // 4 arguments
    sayHello("Mohit", "Rahul", "Rohit", "Johny")
}
  
// using Ellipsis
func sayHello(names ...string) {
    for _, n := range names {
        fmt.Printf("Hello %s\n", n)
    }
}
```

Create a variadic function that sums all it's numbers passed as arguments.

```go
package main

import "fmt"

// Complete the function signature
func sum() int {
  total := 0
    // Your code goes here
}

func main () {
  // Your code goes here
  
}

```

<details>
<summary> Solution: </summary>

```golang
package main

import "fmt"

func sum(numbers ...int) int {
  total := 0
  for _,num := range numbers{
    total += num
  }
  fmt.Println(total)
  return total
}

func main () {
  // Your code goes here
  sum(2,3,4,5,6,7)
}
```

</details>

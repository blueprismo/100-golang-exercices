# Exercise: Recap

This is the final exercise from the basic_exercises block. You will put in practice what you've learned.

In this exercise, we are going to check wether the user has entered a number and if the number is even.

An even number is the one that is divisible by 2 with no remainder.
In order to get the remainder from a division, there is the [remainder/modulo operator](https://en.wikipedia.org/wiki/Modulo) `%`.

For simplicity, we are not going to check if the input is a number or not, we are assuming always a number is entered.

- Define an `int32` variable called `number`
- Create a function to check if a number is even, the function will be called `isEven`
- The function will print to the standard output if the number is even or odd.
- That function will return a `bool` value.

```golang
package main

import "fmt"

func main () {
  // Here goes your code
  fmt.Printf("...")
}
```

<details>
<summary> Solution: </summary>

```golang
package main

import "fmt"

func main () {
  var number int32
  fmt.Println("Enter a number: ")
  fmt.Scanln(&number)

  isEven(number)
}

func isEven(number int32) bool {
  if (number % 2 == 0) {
    fmt.Println("is even")
    return true
  }  else { 
    fmt.Println("is odd")
    return false
  }
}
```

</details>

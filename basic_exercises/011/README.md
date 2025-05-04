# Exercise: Range and Blank identifier

Sometimes we don't need all the variables that are returned from methods or other functions.
In those cases we have a variable called "blank identifier", which is used when we don't want to unnecerarily populate a variable.

The blank identifier is represented with an underscore `_` and it acts as a placeholder for the variable that we do not need.

A simple example in the range loop, we don't care about the values this time but we only want to know the index:

```golang
var arr = [4]int{1,5,6,10}

  for index, _ := range arr {
    fmt.Print(index, "\n")
  }

```

A very common use case for the blank identifier is in exception handling, when we just want to know if a variable returned an error code but not it's content or value, a blank identifier is often used.

Exercise:

- use range to print the values of the array
- use the blank identifier variable to supress the index, DON'T show the index, just the values of the array!

```golang
package main

import "fmt"

func main () {
  // initialized array of 10 int values [1..10]
  var arr = [10]int{1,2,3,4,5,6,7,8,9,10}
  // Here goes your code

}

```

<details>
<summary> Solution: </summary>

```golang

package main

import "fmt"

func main () {

  var arr = [10]int{1,2,3,4,5,6,7,8,9,10}
  
  for _, value := range arr {
    fmt.Print(value, "\n")
  }
}
```

</details>

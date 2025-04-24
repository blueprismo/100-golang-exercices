# Exercise: Pointers

Print the memory address of a variable

```golang
package main

import "fmt"

func main () {
  var x int = 5
  // Your code goes here
  
}

```

<details>
<summary> Solution: </summary>

```golang
package main

import "fmt"

func main () {
  var x int = 5
  // Your code goes here
  fmt.Println(&x)
}
```

</details>

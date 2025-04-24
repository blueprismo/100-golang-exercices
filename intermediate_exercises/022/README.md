# Exercise: Random

Create a program that mimics a dice roll (a 6-face dice)

- A dice roll always returns one number between 1 and 6

Reference:
https://pkg.go.dev/math/rand

```golang
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main () {
	// Your code goes here
	
}
```

<details>
<summary> Solution: </summary>

```golang
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func random(min int, max int) int {
    return rand.Intn(max-min) + min
}

func main() {
    rand.Seed(time.Now().UnixNano())
    randomNum := random(1, 7)
    fmt.Printf("Rolled dice: %d\n", randomNum)
}

```

</details>

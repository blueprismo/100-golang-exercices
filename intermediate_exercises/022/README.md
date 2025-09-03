# Exercise: Random

## Rand package

The rand package in Go helps generate random (or more accurately, pseudo-random) numbers.

To generate pseudo-random numbers, a computer requires a seed, which introduces entropy into the process. If the same seed value is used on every run, the same sequence of pseudo-random numbers will be produced. To generate a different sequence each time, the seed value must be updated.

It is common to see the use of `time.Now().Unix()` as a seed because the current time changes constantly and is unpredictable from one program run to another. This makes it a good source of entropy.

### Why do we need a seed?

Since pseudo-random number generators always produce the same sequence from the same seed, using something dynamic like the current time ensures a different sequence on each execution.

Exercise: Create a program that mimics a dice roll (a 6-face dice)

- A dice roll always returns one number between 1 and 6

Reference:
[https://pkg.go.dev/math/rand](https://pkg.go.dev/math/rand)

```go
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

```go
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

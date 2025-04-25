# Exercise: Random II

Generate a random number from te range [-50, +50]

- seed it with the current time in epoch

To generate a random number within a specific range, you calculate the range width by subtracting the minimum from the maximum, then add the minimum as an offset. This shifts the result into the desired range.

```golang
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {

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
    randomNum := random(-50, +51)
    fmt.Printf("Random number: %d\n", randomNum)
}

```

</details>

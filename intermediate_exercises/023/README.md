# Exercise: Random II

Randomness is not specific to generating a single value, you can also generate a value within a range, a random password containing specific characters, etc.

## Randomness within a range

To generate a random number within a specific range, you calculate the range size by subtracting the minimum from the maximum, then add `1` as an offset. This shifts the result into the desired range.

The function for that is called [rand.Intn()](https://pkg.go.dev/math/rand#Intn)

Because this function is exclusive to the upper bound, that is between `0` and `100` the maximum number it can return is `99`.

Suppose you want a number between 5 and 10 (inclusive).

The numbers you want are:
5, 6, 7, 8, 9, 10 — that's 6 numbers.

So the size of the range is:
max - min + 1 → 10 - 5 + 1 = 6

If we would not add the `+1`, the range would not include the upper bound (10 in the example).

Exercise: Generate a random number from te range [-50, +50]

- seed it with the current time in epoch

```go
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
    return rand.Intn(max - min + 1) + min
}

func main() {
    rand.Seed(time.Now().UnixNano())
    randomNum := random(-50, +50)
    fmt.Printf("Random number: %d\n", randomNum)
}

```

</details>

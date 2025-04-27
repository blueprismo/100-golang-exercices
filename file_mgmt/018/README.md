# Exercise: Write a list of 5 countries to a file

- Tip: use the "os" package

```golang
package main

import "os"


func main () {
    // Here goes your code

}
```

<details>
<summary> Solution: </summary>

```golang
package main

import "os"

func main () {
  // Here goes your code
  file, err := os.Create("create-file.txt")

  if err != nil {
    return
  }

  defer file.Close()

  file.WriteString("Germany\nFrance\nUSA\nSpain\nUK\n")
}
```

</details>

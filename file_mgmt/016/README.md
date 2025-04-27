# Exercise: Check if a file exists

- use the "os" package

```golang
package main

// import ""

func main () {

}
```

<details>
<summary> Solution: </summary>

```golang
package main

import "fmt"
import "os"

func main () {
	// Here goes your code
	if _, err := os.Stat("file-exists.go"); err == nil {
		fmt.Printf("File exists\n");  
	  } else {
		fmt.Printf("File does not exist\n");  
	  }
}
```

</details>

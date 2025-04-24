# Exercise: Read a file

- Tip: use the "io/ioutil" package
- You should run this code where the read file is, or reference it!
- First open the file, then read it thanks to a function call ReadAll()

```golang
func main () {
  // Open the file
	file, err := os.Open()
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the entire file content using io.ReadAll

	// Print the content

}

```

<details>
<summary> Solution: </summary>

```golang
package main

import "fmt"
import "io"
import "os"

func main () {
  // Open the file
	file, err := os.Open("read.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the entire file content using io.ReadAll
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the content
	fmt.Println(string(content))

}
```

</details>

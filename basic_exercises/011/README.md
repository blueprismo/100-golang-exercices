# Exercise: For loops

- Iterate through the array and add two (+2) to each value

```golang

package main

import "fmt"

func main () {
	// initialized array of 10 int values [0..9]
	var arr = [10]int{0,1,2,3,4,5,6,7,8,9}
	// Here goes your code

}
```

<details>
<summary> Solution: </summary>

```golang
package main

import "fmt"

func main () {
	
	var arr = [10]int{0,1,2,3,4,5,6,7,8,9}
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] + 2
		fmt.Println(arr[i])
	}
}
```

</details>

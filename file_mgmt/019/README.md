# Exercise: Rename a file from name1 to name2

- Tip: use the "os" package

```golang
package main

import "os"


func main () {
  var src, dest string
  // Here goes your code

}
```

<details>
<summary> Solution: </summary>

```golang
package main

import "os"


func main () {
	var src, dest string
	// Here goes your code
	
  src = "name1.txt"
  dest = "name2.txt"

  os.Rename(src, dest)
}
```

</details>

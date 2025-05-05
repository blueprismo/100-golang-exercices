# Exercise: Rename a file

This exercise is a simple one, you deserve it!
the "os" package also has a function to [rename](https://pkg.go.dev/os#Rename) a file (as long as you have permissions for it, and that the file exists)

The `os.Rename` function signature:

```go
func Rename(oldpath, newpath string) error
```

Exercise: Rename an existing file from an old name to a new name.

```golang
package main

import "os"


func main () {
  var old_name, new_name string
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

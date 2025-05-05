# Exercise: Check if a file exists

The standard library of golang also provides a package for operative system functionality.

The package is called [os](https://pkg.go.dev/os) and it has a wide range of options for us. It is usually used for file system management.

## Paths

It could be that you are running your code on a linux machine, a Windows one, etc.
It is nice to know how the filesystem is structured. Usually filesystems have a hyerarchy of folders, usually the top-level in linux is the root folder `/` and in Windows it is a convention to use the `C:\\` drive as the root folder.

When you execute your code, you can reference the files with the absolute path on your filesystem or the relative one.

### What is an absolute path?

An absolute path is showing all the filesystem hierarchy with the file we are trying to reference. A techy parallel would be [breadcrumb navigation](https://en.wikipedia.org/wiki/Breadcrumb_navigation).
An example could be `/home/username/app/main.go` for a linux filesystem or `C:\Users\JohnDoe\Documents\MyFile.txt` in a Windows system.

### What is a relative path?

A relative path is only referencing the file we are trying to refer **from the current context** (that is, from the current directory).
Imagine we spawn a terminal in the `/home/username` directory, then our relative path to the file would be `./app/main.go`
In a Windows system, if we spawned a powershell in `C:\Users\JohnDoe` the relative path to Myfile.txt would be `.\Documents\MyFile.txt`.

An important concept with relative paths is that the single dot `.`, represents the current directory, and the double dot `..` represents the parent directory.

Folders also have their own permissions and special attributes, but we're not going to expand on the topic here (as it is huge!).

## Function signature

When you are visiting the documentation of some packages, you will usually find it's function description and the function signature.
When you create a function, you are effectively creating a function signature aswell!
A function signature defines the structure and behaviour of a function, when you write a function, you also write it's signature, let's take the example of the `os.Open()` function signature:

```go
func Open(name string) (*File, error)
// Syntax:
// func <func name>(<argument/s> <argument/s_type>, ...) (return type/s)
```

The Open function will take a string as it's only argument, and will return both the pointer to a File and an error interface.

We can see the `os.Open()` in action:

```go
file, err := os.Open("file.go") // uses the os.Open() function
if err != nil {
  log.Fatal(err)
}
```

Now that you know function signatures, you can quickly go over the documentation of packages and use them!

Exercise: Check if a file exists
Tip: use the [os.Stat()](https://pkg.go.dev/os#Stat) function

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

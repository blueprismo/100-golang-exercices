# Exercise 79: Introduction to Gin Framework

## What is Gin?

Gin is a popular HTTP web framework for Go. It provides a more convenient API than the standard `net/http` package, with features like middleware, JSON binding, and easier routing.

## Why Use Gin?

- **Faster development**: Less boilerplate code
- **Built-in features**: JSON binding, validation, middleware
- **Performance**: Very fast HTTP router
- **Familiar API**: Similar to Express.js (Node.js)

## Installing Gin

```bash
go mod init your-project
go get -u github.com/gin-gonic/gin
```

## Basic Gin Server

```go
import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello World",
        })
    })
    
    r.Run(":8080")
}
```

## Gin vs net/http

**Standard net/http**:
```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}
```

**With Gin**:
```go
func handler(c *gin.Context) {
    c.JSON(200, data)
}
```

## Your Task

Create your first Gin web server:
1. Install Gin framework
2. Create a basic server with Gin
3. Set up a simple route that returns JSON
4. Compare the simplicity to standard net/http

This introduces modern Go web development with frameworks.
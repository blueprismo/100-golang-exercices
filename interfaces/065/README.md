# Exercise 65: Empty Interface

## What is interface{}?

The empty interface `interface{}` specifies zero methods, which means any type satisfies it. It's Go's way of representing "any type" - similar to `Object` in Java or `any` in TypeScript.

## Why Empty Interface?

Since every type implements zero methods, every type implements the empty interface:

```go
var anything interface{}
anything = 42         // int
anything = "hello"    // string  
anything = []int{1,2} // slice
anything = true       // bool
```

## Common Use Cases

- **Generic data structures**: When you need to store different types
- **JSON unmarshaling**: Unknown data structures
- **Function parameters**: Accept any type
- **Reflection**: Runtime type inspection

## Type Safety Warning

Empty interfaces sacrifice type safety for flexibility. Use them sparingly and prefer concrete types when possible.

## Your Task

Create a `human` type that uses `map[string]interface{}` to store various data types:
- Keys are strings (field names)
- Values can be any type (name, age, etc.)

This demonstrates how empty interfaces enable flexible data structures at the cost of type safety.
# Exercise 68: Type Switch

## What is a Type Switch?

A type switch is a special form of switch statement that allows you to determine and handle the concrete type of an interface{} value. It's more elegant than multiple type assertions.

## Type Switch Syntax

```go
func processValue(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

## Key Points

- **`i.(type)`**: Special syntax only valid in type switches
- **Variable assignment**: `v` gets the concrete value
- **Type-specific handling**: Each case handles a different type
- **Default case**: Handles unexpected types

## Why Use Type Switches?

- **Cleaner than multiple type assertions**
- **More readable than long if-else chains**
- **Performance**: More efficient than repeated assertions
- **Comprehensive**: Easy to handle many types

## Your Task

Create a function that:
1. Accepts an empty interface as parameter
2. Uses a type switch with `i.(type)` syntax
3. Handles different types (int, string, etc.) with specific logic
4. Includes a default case for unknown types

This demonstrates the most elegant way to handle multiple types with empty interfaces.
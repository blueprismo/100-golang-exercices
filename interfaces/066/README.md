# Exercise 66: Type Assertions with Empty Interfaces

## The Problem with Empty Interfaces

When you store values in `interface{}`, you lose type information. To use the values, you need **type assertions** to convert back to concrete types.

## Type Assertion Syntax

```go
// Basic type assertion
value := someInterface.(int)

// Safe type assertion with ok check  
value, ok := someInterface.(int)
if !ok {
    fmt.Println("Not an int!")
}
```

## Why Type Assertions Are Needed

```go
var x interface{} = 42
// x + 1  // ERROR: can't add to interface{}
y := x.(int)  // Type assertion
result := y + 1  // OK: now it's an int
```

## Safe vs Unsafe Assertions

**Unsafe** (panics if wrong type):
```go
age := person["age"].(int)  // Panic if not int
```

**Safe** (returns bool):
```go
if age, ok := person["age"].(int); ok {
    fmt.Println("Age:", age + 1)
} else {
    fmt.Println("Age is not an integer")
}
```

## Your Task

Work with a `human` map containing mixed types. Practice:
1. Storing different types (string, int) in the map
2. Using type assertions to retrieve and modify values
3. Handling the case where you need to increment an age value

This demonstrates both the power and complexity of working with empty interfaces.
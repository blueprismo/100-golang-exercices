# Exercise 57: Context Chaining and Immutability

## Context Chaining

Context values can be chained - you can create new contexts that inherit from parent contexts and add or modify values. This creates a chain of contexts, each building on the previous one.

## How Context Chaining Works

```go
ctx1 := context.Background()
ctx2 := context.WithValue(ctx1, "key1", "value1")  
ctx3 := context.WithValue(ctx2, "key2", "value2")
ctx4 := context.WithValue(ctx3, "key1", "newvalue") // Shadows "key1" from ctx2
```

## Context Immutability

**Important**: Contexts are immutable! When you call `context.WithValue()`, you don't modify the existing context - you create a new one.

```go
originalCtx := context.WithValue(context.Background(), "name", "John")
newCtx := context.WithValue(originalCtx, "name", "Jane")

// originalCtx still has "John"
// newCtx has "Jane"  
// They are separate contexts!
```

## Value Shadowing

When you add a value with the same key to a child context, it "shadows" (hides) the parent's value:

```go
parent := context.WithValue(context.Background(), "user", "Alice")
child := context.WithValue(parent, "user", "Bob")

fmt.Println(parent.Value("user")) // "Alice"
fmt.Println(child.Value("user"))  // "Bob" (shadows Alice)
```

## Your Task

1. **Complete the `doAnother` function** that prints the same key value
2. **In `doSomething`, create a new context** using `context.WithValue()` with the same key but different value
3. **Call `doAnother` with the new context**
4. **Observe the behavior**: Does the value change in the child context?
5. **Test immutability**: Check if the original context's value changed

## Key Learning Points

- Contexts form a tree structure
- Child contexts can shadow parent values
- Original contexts remain unchanged (immutable)
- Each `WithValue` call creates a new context
# Exercise 63: Implicit Interface Implementation

## Go's Implicit Interface Implementation

Go uses **implicit interface implementation** - you don't explicitly declare that a type implements an interface. If a type has all the methods that an interface requires, it automatically implements that interface.

## No "implements" Keyword

Unlike other languages, Go has no `implements` keyword:

```go
// Java/C# style (NOT Go):
// class Rectangle implements Geometry { ... }

// Go style - just implement the methods:
type Rectangle struct { ... }
func (r Rectangle) area() float64 { ... }  // Now Rectangle implements Geometry!
```

## How Implicit Implementation Works

1. **Interface defines contract**: "Types must have these methods"
2. **Type provides methods**: Methods with exact same signatures
3. **Go compiler checks**: "Does this type satisfy the interface?"
4. **Automatic implementation**: No explicit declaration needed

## Method Receivers

When implementing interface methods, you typically use method receivers:

```go
type rect struct {
    width, height float64
}

// Method with receiver - implements interface method
func (r rect) area() float64 {
    return r.width * r.height
}
```

## Benefits of Implicit Implementation

- **Flexibility**: Existing types can implement new interfaces
- **Decoupling**: Interface and implementation can be in different packages
- **Simplicity**: Less boilerplate code
- **Retroactive implementation**: Add interface support to existing types

## Your Task

1. **Implement the `area()` method** for the `rect` type
2. **Use a method receiver**: `func (r rect) area() float64`
3. **Calculate rectangle area**: width × height
4. **Create and test** a rectangle instance

## Key Learning

By implementing the `area()` method, your `rect` type automatically satisfies the `geometry` interface - no explicit declaration needed!
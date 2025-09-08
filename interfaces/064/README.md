# Exercise 64: Extending Interfaces

## Adding Methods to Interfaces

Interfaces can have multiple method signatures. When you add a method to an interface, all implementing types must provide that method to remain compatible.

## Interface Evolution

```go
// Version 1: Simple interface
type Shape interface {
    Area() float64
}

// Version 2: Extended interface  
type Shape interface {
    Area() float64
    Perimeter() float64  // New method added
}
```

## Multiple Method Implementation

Types must implement ALL methods in an interface:

```go
type Rectangle struct {
    width, height float64
}

// Must implement both methods
func (r Rectangle) Area() float64 { /* ... */ }
func (r Rectangle) Perimeter() float64 { /* ... */ }
```

## Rectangle Perimeter Formula

For a rectangle:
- **Perimeter = 2 × width + 2 × height**
- **Or: 2 × (width + height)**

## Your Task

1. **Add `perim()` method** to the geometry interface
2. **Implement `perim()` for rect type** using the perimeter formula
3. **Test both methods** - area and perimeter

## Key Concept

When you extend an interface, existing implementing types must be updated to include the new methods, or they'll no longer satisfy the interface.
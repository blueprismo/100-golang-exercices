# Exercise 54: Introduction to Context

## What is Context?

The `context` package in Go provides a way to carry deadlines, cancellation signals, and request-scoped values across API boundaries and between processes. Context is fundamental for building robust, cancellable operations in Go.

## Why Context is Important

Context solves several critical problems:
- **Cancellation**: Stop operations when they're no longer needed
- **Timeouts**: Prevent operations from running too long  
- **Request tracing**: Pass request-specific data through call chains
- **Resource cleanup**: Ensure proper cleanup when operations are cancelled

## The Context Interface

`context.Context` is an interface with four methods:
```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error  
    Value(key interface{}) interface{}
}
```

## Getting Started: context.TODO()

`context.TODO()` returns an empty context. It's used when:
- You're not sure which context to use
- You're in the process of adding context support
- You need a placeholder during development

```go
ctx := context.TODO()
```

## Context as First Parameter

By Go convention, context is typically the first parameter of functions:
```go
func DoWork(ctx context.Context, data string) error {
    // Function implementation
}
```

## Key Concepts

- **Context is immutable**: You can't modify a context, only derive new ones
- **Context is safe for concurrent use**: Multiple goroutines can use the same context
- **Context forms a tree**: Child contexts inherit from parent contexts

## Your Task

1. **Import the context package**
2. **Create a function `doSomething`** that accepts `context.Context` as a parameter
3. **In main, create a context** using `context.TODO()`
4. **Call your function** with the context

This is your first step into the world of Go contexts - understanding how to pass them between functions.
# Exercise 32: Introduction to Goroutines

## What are Goroutines?

Goroutines are lightweight threads managed by the Go runtime. They are one of Go's most powerful features for concurrent programming. Unlike operating system threads, goroutines are much more lightweight - you can create thousands or millions of them with minimal overhead.

## Key Characteristics of Goroutines

- **Lightweight**: Start with small stack (around 2KB) that grows as needed
- **Managed**: Go runtime handles scheduling onto OS threads
- **Concurrent**: Multiple goroutines can run simultaneously
- **Easy to create**: Just use the `go` keyword before a function call

## Creating Goroutines

To create a goroutine, simply prefix any function call with the `go` keyword:

```go
go functionName() // This function now runs concurrently
```

## Important Concepts

**Main Goroutine**: Every Go program starts with one goroutine running the `main()` function. When the main goroutine terminates, the entire program terminates, even if other goroutines are still running.

**Scheduling**: The Go runtime automatically manages when and where goroutines run. You don't control which OS thread they run on.

**Why Use time.Sleep()**: In early goroutine examples, you'll often see `time.Sleep()` to give goroutines time to execute before the main function ends. This is a temporary solution - later you'll learn better synchronization methods like channels and WaitGroups.

## Your Task

Look at the `main.go` file and complete the exercise. The goal is to understand how to:
1. Create goroutines using the `go` keyword
2. Observe concurrent execution
3. Understand why we need to wait for goroutines to complete

## Expected Behavior

When you run your completed program, you should see output from both function calls, but the order might be different each time due to concurrent execution.

## Hint

Remember that the `go` keyword turns any function call into a goroutine. The function signature doesn't change - you just add `go` before calling it.
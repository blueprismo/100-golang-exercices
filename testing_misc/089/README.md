# Exercise 89: Mutexes and WaitGroups

## Concurrency Synchronization

When multiple goroutines access shared data, you need synchronization mechanisms to prevent race conditions and ensure data consistency. Go provides mutexes for exclusive access and WaitGroups for coordination.

## Race Conditions

Race conditions occur when:
- Multiple goroutines access the same variable
- At least one access is a write operation  
- No synchronization mechanism prevents concurrent access

## Mutexes (sync.Mutex)

Mutexes provide mutual exclusion - only one goroutine can hold the lock at a time:

```go
type SafeCounter struct {
    mu sync.Mutex
    n  int
}

func (c *SafeCounter) Inc() {
    c.mu.Lock()   // Acquire lock
    c.n++         // Critical section
    c.mu.Unlock() // Release lock
}
```

## WaitGroups (sync.WaitGroup)

WaitGroups coordinate multiple goroutines:
- `Add(n)`: Add n goroutines to wait for
- `Done()`: Signal one goroutine completed
- `Wait()`: Block until all goroutines call Done()

```go
var wg sync.WaitGroup
wg.Add(5) // Wait for 5 goroutines

for i := 0; i < 5; i++ {
    go func() {
        defer wg.Done() // Signal completion
        // Do work here
    }()
}
wg.Wait() // Wait for all to finish
```

## Critical Sections

Code that accesses shared resources must be protected:
- Only one goroutine can execute critical section at a time
- Use `defer mutex.Unlock()` to ensure unlock happens
- Keep critical sections as small as possible

## Your Task

Look at the main.go file and complete the exercise by:
1. Implementing mutex locking in the Inc() method
2. Using defer to ensure proper unlocking
3. Setting up WaitGroup to coordinate goroutines
4. Understanding how synchronization prevents race conditions

This exercise demonstrates essential concurrency patterns for safe shared data access in Go programs.
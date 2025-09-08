# Exercise 86: Benchmark Testing

## What is Benchmarking?

Benchmarking measures the performance of your code by running it many times and measuring execution time, memory allocations, and other performance metrics. Go's testing package includes built-in benchmarking support.

## Benchmark Functions

Benchmark functions must:
- Start with `Benchmark` prefix
- Accept `*testing.B` parameter
- Use `b.N` in a loop to run the code being tested
- Have signature: `func BenchmarkXxx(b *testing.B)`

## The b.N Loop

The testing framework automatically adjusts `b.N` to get reliable timing measurements:
- Starts with small N, increases until benchmark runs long enough
- Typically runs for at least 1 second
- Reports operations per second (ns/op, μs/op, etc.)

## Example Benchmark Structure

```go
func BenchmarkExample(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // Code being benchmarked
        result := ExpensiveFunction()
        _ = result // Prevent optimization
    }
}
```

## Running Benchmarks

```bash
go test -bench .           # Run all benchmarks
go test -bench BenchmarkXxx # Run specific benchmark
go test -bench . -benchmem  # Include memory stats
```

## Benchmark Output

```
BenchmarkExample-8    1000000    1234 ns/op    456 B/op    7 allocs/op
```
- Function name and GOMAXPROCS
- Number of iterations
- Nanoseconds per operation
- Bytes allocated per operation
- Allocations per operation

## Your Task

Look at the main_test.go file and complete the exercise by:
1. Creating a benchmark function with proper signature
2. Using b.N in a loop to measure performance
3. Understanding how Go's benchmark system works

This exercise introduces performance testing and helps you identify bottlenecks in your code.
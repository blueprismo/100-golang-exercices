# Exercise 84: Test Failures

## Understanding Test Failures

Testing isn't just about passing tests - understanding what happens when tests fail is equally important. Failed tests provide valuable feedback about bugs, logic errors, and edge cases in your code.

## When Tests Fail

A test fails when:
- The actual result doesn't match the expected result
- An assertion fails (using t.Error, t.Fail, etc.)
- A panic occurs during test execution
- The function behavior differs from expectations

## Analyzing Test Output

When a test fails, Go provides:
- Which test function failed
- Line number of the failure
- Error message you provided
- Stack trace if applicable

```bash
--- FAIL: TestExample (0.00s)
    example_test.go:10: Expected "Hello", got "Goodbye"
FAIL
```

## Common Failure Scenarios

- **Logic Errors**: Function returns wrong value
- **Edge Cases**: Boundary conditions not handled
- **State Issues**: Shared state between tests
- **Dependency Problems**: External services unavailable

## Your Task

Look at the main.go file and complete the exercise by:
1. Creating a function that will intentionally cause test failure
2. Understanding how Go reports test failures
3. Learning to interpret test failure messages

This exercise teaches you to recognize and understand test failures, which is crucial for debugging and maintaining code quality.
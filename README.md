# Go Beginner Level Coding Test

This is the beginner developer coding test for [Go](https://golang.org/). Please fork this repository into your own account or download the code and push it into your own public repository.

# Duplicates

Write the code for **Duplicates** function in _duplicates.go_ so the provided unit test works as expected.

## Unit test

Your code should pass the provided unit test successfully.

```go
package duplicates

import (
	"testing"
	"bytes"
)

func TestDuplicates(t *testing.T) {
    origin := []byte {1,3,2,3,24,36,2,5,90,122,15,45,32,24}
	expected := []byte {2,3,24}

	result := Duplicates(origin)

    if !bytes.Equal(result, expected) {
       t.Errorf("Result was incorrect, got: %v, expected: %v.", result, expected)
    }
}
```

**Plus**

> Create an implementation of Table Driven tests for the Duplicates function.

# Fibonacci

Write a function that calculates a given Fibonacci number in the [fibonacci sequence](https://en.wikipedia.org/wiki/Fibonacci_number), this implementation should be recursive.

**Plus**

> Provide unit test implementation.
> Provide a not recursive solution.
> Calculate the maximum n your recursive function can handle.

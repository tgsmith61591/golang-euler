package euler

import (
	"fmt"
	"time"
)

// This allows us to pass around a function call to the TimedExecution method,
// as some euler solutions will be floats or strings as opposed to ints.
type EulerSolution struct{
	Solution string
	Err error
}

// Type alias for `func() *EulerSolution` so we don't have to cart that
// around in the method signatures
type Euler func() *EulerSolution


// Executes a function and prints the time it took to run
func TimedExecution(name string, f Euler) {
	start := time.Now()
	result := f().Solution
	elapsed := time.Since(start)
	fmt.Printf("%s:\t%s (%s)\n", name, result, elapsed)
}

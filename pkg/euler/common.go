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
func TimedExecution(name string, f Euler, expected string) {
	start := time.Now()
	result := f()
	elapsed := time.Since(start)
	
	if result.Err != nil {
		fmt.Printf("%s (FAIL):\t%s (%s)\n", name, result.Err, elapsed)
	} else {
		if result.Solution != expected {
			fmt.Printf("%s (FAIL):\t%s != %s(%s)\n",
					   name, result.Solution, expected, elapsed)
		} else {
			fmt.Printf("%s (PASS):\t%s (%s)\n", name, result.Solution, elapsed)
		}
	}
}

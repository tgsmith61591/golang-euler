package euler

import (
	"fmt"
	"time"
)

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

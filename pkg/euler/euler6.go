package euler

import (
	"strconv"
)

/**
The sum of the squares of the first ten natural numbers is,
   1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,

  (1 + 2 + ... + 10)^2 = 552 = 3025

Hence the difference between the sum of the squares of the first ten natural
numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred
natural numbers and the square of the sum.
 */
func Euler6() *EulerSolution {
	sumOfSq := 0
	sqOfSum := 0
	
	for i := 1; i < 101; i++ {
		sumOfSq += i * i
		sqOfSum += i
	}

	// square this one
	sqOfSum *= sqOfSum
	res := sqOfSum - sumOfSq

	return &EulerSolution{
		Solution: strconv.Itoa(res),
	}
}

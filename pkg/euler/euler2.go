package euler

import (
	"strconv"
)


func Euler2() *EulerSolution {
	a := 1
	b := 2
	var tmp int
	sum := b

	// By considering the terms in the Fibonacci sequence whose values do not
	// exceed four million...
	for b < 4000000 {
		tmp = a + b
		a = b
		b = tmp

		if b % 2 == 0 {
			sum += b
		}
	}

	return &EulerSolution{
		Solution: strconv.Itoa(sum),
	}
}

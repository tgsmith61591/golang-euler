package euler

import (
	"strconv"
)

/**
  The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

  Find the sum of all the primes below two million.
 */
func Euler10() *EulerSolution {
	sieve, err := SieveOfEratosthenes(2000000)
	if err != nil {
		return &EulerSolution{
			Err: err,
		}
	}

	sum := 0
	for _, e := range sieve {
		sum += e
	}

	return &EulerSolution{
		Solution: strconv.Itoa(sum),
	}
}

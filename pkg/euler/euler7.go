package euler

import (
	"strconv"
)

/**
 By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see
 that the 6th prime is 13.

 What is the 10 001st prime number?
 */
func Euler7() *EulerSolution {
	// Not elegant, but generate an arbitrarily high num of primes and find the
	// 10001st
	primes, err := SieveOfEratosthenes(150000)
	if err != nil {
		return &EulerSolution{
			Err: err,
		}
	}

	res := primes[10000]  // remember, 0 indexed..
	return &EulerSolution{
		Solution: strconv.Itoa(res),
	}
}

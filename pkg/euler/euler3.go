package euler

import (
	"strconv"
)

// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

func Euler3() *EulerSolution {
	n := 600851475143

	// get a sieve of eratosthenes up to some arbitrarily high number...
	sieve, err := SieveOfEratosthenes(10000)
	if err != nil {
		return &EulerSolution{
			Err: err,
		}
	}

	var res int
	for i := len(sieve) - 1; i >= 0; i-- {
		if n % sieve[i] == 0 {
			res = sieve[i]
			break
		}
	}

	return &EulerSolution{
		Solution: strconv.Itoa(res),
	}
}

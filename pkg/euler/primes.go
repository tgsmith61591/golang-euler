package euler

import (
	"errors"
)


// Very fast way to generate primes to up n (inclusive)
func SieveOfEratosthenes(n int) ([]int, error) {
	if n < 3 {
		return nil, errors.New("n must be >= 3")
	}

	// Initialize a list to n (inclusive)
	var sieve []int
	var mask []bool
	for i := 0; i < n + 1; i++ {
		sieve = append(sieve, i)
		mask = append(mask, true)  // initialize all to true until marked
	}

	// Mark 0 and 1 as false, as they are not prime technically
	mask[0] = false
	mask[1] = false

	p := 2  // initially, let p equal 2, the smallest prime number

	// enumerate the multiples of p by counting in increments of p from 2p to n
	// and mark them in the list
	search:
	for {
		mask[p] = true
		for i := 2 * p; i < n + 1; i += p {
			mask[i] = false
		}

		// Find the first number greater than p in the list that is not marked.
		// If there is no such number, stop.
		for i := p + 1; i < n + 1; i++ {
			if mask[i] {  // still candidate for being a prime
				p = i
				break
			} else if i == n {
				// nothing matched
				break search
			}
		}
	}

	// Now get the indices that are true and return
	var out []int
	for i, val := range mask {
		if val {
			out = append(out, i)
		}
	}

	return out, nil
}

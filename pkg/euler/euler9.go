package euler

import (
	"strconv"
	"errors"
)

/**
 A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

   a^2 + b^2 = c^2

 For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

 There exists exactly one Pythagorean triplet for which a + b + c = 1000.
 Find the product abc.
 */
func Euler9() *EulerSolution {

	// I don't have a great solution in mind right now, so I'll just
	// brute force it...
	for a := 0; a < 500; a++ {
		for b := a + 1; b < 500; b++ {
			if (a * a + b * b == (1000 - a - b) * (1000 - a - b)) {
				c := (1000 - a - b)
				return &EulerSolution{
					Solution: strconv.Itoa(a * b * c),
				}
			}
		}
	}

	return &EulerSolution{
		Err: errors.New("did not find a solution"),
	}
}

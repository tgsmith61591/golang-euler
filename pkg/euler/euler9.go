package euler

import (
	_ "strconv"
	"errors"
	_ "math"
)

/**
 A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

   a^2 + b^2 = c^2

 For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

 There exists exactly one Pythagorean triplet for which a + b + c = 1000.
 Find the product abc.
 */
func Euler9() *EulerSolution {

	// Euclid's formula for generating triplets:
	//   a = m^2 - n^2
	//   b = 2m*n
	//   c = m^2 + n^2
	// For all m > n > 0
	// TODO:

	return &EulerSolution{
		Err: errors.New("did not find a solution"),
	}
}

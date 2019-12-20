package euler

import (
	"strconv"
)

// 2520 is the smallest number that can be divided by each of the numbers from
// 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the
// numbers from 1 to 20?
func Euler5() *EulerSolution {
	// I am not a very smartboi and can't think of a super clever way to solve
	// this, so I'll just brute force it.
	res := 2520
	for res % 20 != 0 || 
		res % 19 != 0 ||
		res % 18 != 0 ||
		res % 17 != 0 ||
		res % 16 != 0 ||
		res % 15 != 0 ||
		res % 14 != 0 ||
		res % 13 != 0 ||
		res % 12 != 0 ||
		res % 11 != 0 ||
		res % 7 != 0 {
		// This is one little clever hackaroonie. Don't need to increment by 1
		res += 20
	}

	return &EulerSolution{
		Solution: strconv.Itoa(res),
	}
}

package euler

import (
	"strconv"
)


// reverse an integer with no string operations
func reverseInteger(n int) int {
	if n < 10 {
		return n
	}

	out := 0
	for n > 0 {
		remainder := n % 10  // strip off end
		out = out * 10 + remainder  // shift over the output
		n /= 10
	}

	return out
}

func isPalindrome(n int) bool {
	return n == reverseInteger(n)
}

// A palindromic number reads the same both ways. The largest palindrome made
// from the product of two 2-digit numbers is 9009 = 91 × 99
// Find the largest palindrome made from the product of two 3-digit numbers.
func Euler4() *EulerSolution {
	var i int = 999
	var res int = 0
	for i > 99 {
		var j int = i
		for j > 99 {
			prod := i * j
			if isPalindrome(prod) && prod > res {
				res = prod
			}
			j -= 1
		}
		i -= 1
	}

	return &EulerSolution{
		Solution: strconv.Itoa(res),
	}
}

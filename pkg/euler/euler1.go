package euler

import (
	"strconv"
)

func Euler1() *EulerSolution {
	sum := 0
	for i := 0; i < 1000; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	return &EulerSolution{
		Solution: strconv.Itoa(sum),
	}
}

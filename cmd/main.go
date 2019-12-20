package main

import (
	e "github.com/golang-euler/pkg/euler"
)

func main() {
	e.TimedExecution("Euler 1", e.Euler1, `233168`)
	e.TimedExecution("Euler 2", e.Euler2, `4613732`)
	e.TimedExecution("Euler 3", e.Euler3, `6857`)
	e.TimedExecution("Euler 4", e.Euler4, `906609`)
	e.TimedExecution("Euler 5", e.Euler5, `232792560`)
	e.TimedExecution("Euler 6", e.Euler6, `25164150`)
	e.TimedExecution("Euler 7", e.Euler7, `104743`)
	e.TimedExecution("Euler 8", e.Euler8, `23514624000`)
	e.TimedExecution("Euler 9", e.Euler9, `TODO`)
}

package triangle

import (
	"math"
)

const testVersion = 3

type Kind int

var (
	NaT Kind = 1
	Equ Kind = 2
	Iso Kind = 3
	Sca Kind = 4
)

func KindFromSides(a, b, c float64) Kind {
	l1 := a + b
	l2 := a + c
	l3 := b + c

	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}

	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return NaT
	}

	if math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1) {
		return NaT
	}

	if !(l1 >= c && l2 >= b && l3 >= a) || ((a + b + c) == 0) {
		return NaT
	}

	if a == b && a == c && b == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}

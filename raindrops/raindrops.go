package raindrops

import (
	"strconv"
)

const testVersion = 3

func Convert(n int) string {
	x, y, z := n%3, n%5, n%7
	if x != 0 && y != 0 && z != 0 {
		return strconv.Itoa(n)
	}

	out := ""
	if x == 0 {
		out += "Pling"
	}

	if y == 0 {
		out += "Plang"
	}

	if z == 0 {
		out += "Plong"
	}
	return out
}

// Don't forget the test program has a benchmark too.
// How fast does your Convert convert?

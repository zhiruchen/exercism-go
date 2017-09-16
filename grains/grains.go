package grains

import (
	"errors"
)

const testVersion = 1

func Square(n int) (uint64, error) {
	if n <= 0 || n == 65 {
		return 0, errors.New("n is wrong")
	}

	return 1 << uint(n-1), nil
}

func Total() uint64 {
	var sum uint64

	var i uint = 1
	for i <= 64 {
		sum += 1 << (i - 1)
		i++
	}
	return sum
}

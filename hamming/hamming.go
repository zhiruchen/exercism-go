package hamming

import (
	"errors"
)

const testVersion = 6

func Distance(a, b string) (int, error) {
	if a == "" && b == "" {
		return 0, nil
	}

	as := []rune(a)
	bs := []rune(b)

	if len(as) != len(bs) {
		return -1, errors.New("disallow")
	}

	count := 0
	for i := 0; i < len(as); i++ {
		if as[i] != bs[i] {
			count++
		}
	}
	return count, nil
}

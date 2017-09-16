package queenattack

import (
	"errors"
	"math"
)

const testVersion = 2

var err = errors.New("can not attack")

func CanQueenAttack(a, b string) (bool, error) {
	if a == b || a == "" || b == "" {
		return false, err
	}

	aa := []rune(a)
	bb := []rune(b)
	if len(aa) > 2 || len(bb) > 2 {
		return false, err
	}

	if aa[1] < '1' || aa[1] > '8' || bb[1] < '1' || bb[1] > '8' {
		return false, err
	}

	if aa[0] > 'h' || bb[0] > 'h' {
		return false, err
	}

	if aa[0] == bb[0] || aa[1] == bb[1] {
		return true, nil
	}

	d1 := math.Abs(float64(aa[0] - bb[0]))
	d2 := math.Abs(float64(aa[1] - bb[1]))

	return d1 == d2, nil
}

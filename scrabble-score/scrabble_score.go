package scrabble

import (
	"fmt"
)

const testVersion = 5

var scoreTable = map[string]int{
	"A": 1,
	"B": 3,
	"E": 1,
	"I": 1,
	"O": 1,
	"U": 1,
	"L": 1,
	"N": 1,
	"R": 1,
	"S": 1,
	"T": 1,
	"D": 2,
	"G": 2,
	"C": 3,
	"M": 3,
	"P": 3,
	"F": 4,
	"H": 4,
	"V": 4,
	"W": 4,
	"Y": 4,
	"K": 5,
	"J": 8,
	"X": 8,
	"Q": 10,
	"Z": 10,
}

func Score(s string) int {
	score := 0
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			letter = letter - 32
		}

		key := fmt.Sprintf("%c", letter)
		v, _ := scoreTable[key]
		score += v
	}
	return score
}

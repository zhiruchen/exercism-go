package anagram

import (
	"strings"
)

const testVersion = 2

func Detect(subject string, candidates []string) []string {
	subjectMap := map[rune]int{}

	for _, r := range subject {
		if r >= 'A' && r <= 'Z' {
			r = r + 32
		}
		subjectMap[r] = subjectMap[r] + 1
	}

	anagramMap := map[string]bool{}
	newMap := map[rune]int{}

OUTER:
	for _, s := range candidates {
		if len(subject) != len(s) {
			continue OUTER
		}

		if s == subject || strings.ToLower(s) == strings.ToLower(subject) {
			continue OUTER
		}

		newMap = copyMap(subjectMap)
		for _, c := range s {
			if c >= 'A' && c <= 'Z' {
				c = c + 32
			}
			if v, ok := newMap[c]; ok {
				newMap[c] = v - 1
			}
		}
		for _, num := range newMap {
			if num != 0 {
				newMap = map[rune]int{}
				continue OUTER
			}
		}

		anagramMap[s] = true
		newMap = map[rune]int{}
	}

	anagramList := []string{}
	for k := range anagramMap {
		anagramList = append(anagramList, k)
	}
	return anagramList
}

func copyMap(m map[rune]int) map[rune]int {
	var copy = map[rune]int{}

	for k, v := range m {
		copy[k] = v
	}

	return copy
}

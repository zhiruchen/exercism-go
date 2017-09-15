package series

import (
	"log"
)

const testVersion = 2

func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}

	ss := []rune(s)
	sLen := len(s)

	subList := []string{}

	for i := 0; i < len(s); i++ {
		sub := []rune{}

		if i+n > sLen {
			break
		}

		for j := i; j < i+n; j++ {
			sub = append(sub, ss[j])
		}

		subList = append(subList, string(sub))
	}
	return subList
}

func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		return ""
	}

	ss := []rune(s)
	sub := []rune{}

	for j := 0; j < n; j++ {
		sub = append(sub, ss[j])
	}
	return string(sub)
}

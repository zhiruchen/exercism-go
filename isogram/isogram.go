package isogram

import (
	"fmt"
)

const testVersion = 1

func IsIsogram(word string) bool {
	var m = make(map[string]bool)

	for _, s := range word {
		if s >= 'A' && s <= 'Z' {
			s = s + 32
		}
		ss := fmt.Sprintf("%c", s)
		if ss == " " || ss == "-" {
			continue
		}

		_, ok := m[ss]
		if ok {
			return false
		}
		m[ss] = true
	}
	return true
}

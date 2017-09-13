package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 3

func Abbreviate(s string) string {
	ss := strings.Split(s, " ")

	rep := ""
	for _, x := range ss {
		ws := []string{}
		if strings.Contains(x, "-") {
			ws = strings.Split(x, "-")
		} else {
			ws = append(ws, x)
		}

		for _, w := range ws {
			for _, r := range w {
				rep += string(unicode.ToUpper(r))
				break
			}
		}
	}
	return rep
}

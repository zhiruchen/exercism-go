package etl

import "strings"

const testVersion = 1

func Transform(input given) expectation {
	var m = make(expectation)

	for num, letters := range input {
		for _, s := range letters {
			m[strings.ToLower(s)] = num
		}
	}

	return m
}

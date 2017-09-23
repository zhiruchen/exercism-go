package sieve

import (
	"math"
)

const testVersion = 1

func Sieve(limit int) []int {
	intSlice := []bool{false, false}
	for i := 2; i < limit+1; i++ {
		intSlice = append(intSlice, true)
	}

	for i := 2; i <= int(math.Sqrt(float64(limit))); i++ {
		if intSlice[i] {
			j := i * i
			for j <= limit {
				intSlice[j] = false
				j += i
			}
		}
	}

	result := []int{}
	for i := 2; i < limit+1; i++ {
		if intSlice[i] {
			result = append(result, i)
		}
	}
	return result
}

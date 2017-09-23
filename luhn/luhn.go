package luhn

import (
	"strings"
)

const testVersion = 2

var numMap = map[rune]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func Valid(number string) bool {
	if len(number) <= 1 {
		return false
	}

	number = strings.Trim(number, " ")
	if len(number) <= 1 {
		return false
	}

	number = strings.Replace(number, " ", "", -1)
	numberList := []int{}

	for _, n := range number {
		if !checkNumber(n) {
			return false
		}

		v, ok := numMap[n]
		if ok {
			numberList = append(numberList, v)
		}
	}

	nIndex := len(numberList) - 2
	for nIndex >= 0 {
		numberList[nIndex] *= 2
		if numberList[nIndex] > 9 {
			numberList[nIndex] -= 9
		}
		nIndex -= 2
	}

	sum := 0
	for _, n := range numberList {
		sum += n
	}

	return sum%10 == 0
}

func checkNumber(n rune) bool {
	return n >= '0' && n <= '9'
}

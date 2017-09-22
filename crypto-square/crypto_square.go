package cryptosquare

import (
	"math"
)

const testVersion = 2

func Encode(pt string) string {
	if pt == "" {
		return ""
	}

	normalizedList := normalize(pt)
	letterLen := len(normalizedList)

	rowsLen := int(math.Floor(math.Sqrt(float64(letterLen))))
	var colLen = 0
	if rowsLen*rowsLen == letterLen {
		colLen = rowsLen
	} else {
		colLen = rowsLen + 1
	}
	if colLen*rowsLen < letterLen {
		rowsLen++
	}

	letterMatrix := make([][]rune, rowsLen)
	for i := range letterMatrix {
		letterMatrix[i] = make([]rune, colLen)
	}

	var letterIndex = 0
	for _, row := range letterMatrix {
		for j := 0; j < len(row); j++ {
			if letterIndex >= letterLen {
				row[j] = ' '
				letterIndex++
				continue
			}
			row[j] = normalizedList[letterIndex]
			letterIndex++
		}
	}

	var resultList = []rune{}
	var rowIndex = 0
	for i := 0; i < colLen; i++ {
		for rowIndex < rowsLen {
			v := letterMatrix[rowIndex][i]
			if v == ' ' {
				rowIndex++
				continue
			}
			resultList = append(resultList, v)
			rowIndex++
		}
		if i == colLen-1 {
			continue
		}
		resultList = append(resultList, ' ')
		rowIndex = 0
	}

	return string(resultList)
}

func normalize(s string) []rune {
	runeLst := []rune{}

	for _, x := range s {
		if (x >= '1' && x <= '9') || (x >= 'a' && x <= 'z') || (x >= 'A' && x <= 'Z') {
			if x >= 'A' && x <= 'Z' {
				x = x + 32
			}
			runeLst = append(runeLst, x)
		}
	}
	return runeLst
}

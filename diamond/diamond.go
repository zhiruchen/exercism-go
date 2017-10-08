package diamond

import (
	"errors"
	"strings"
)

const testVersion = 1

func Gen(letter byte) (string, error) {
	letter1 := rune(letter)
	if !(letter1 >= 'A' && letter1 <= 'Z') {
		return "", errors.New("letter out of range")
	}

	if letter1 == 'A' {
		return "A\n", nil
	}

	if letter1 > 'A' {
		result := []string{}

		dis := int(letter1 - 'A')
		for i := 0; i <= dis; i++ {
			row := ""
			for j := 1; j <= dis-i; j++ {
				row += " "
			}

			if i == 0 {
				row += string('A')
			} else {
				row += string('A' + i)
				for x := 1; x <= (2*i)-1; x++ {
					row += " "
				}
				row += string('A' + i)
			}

			for j := 1; j <= dis-i; j++ {
				row += " "
			}

			result = append(result, row)
		}

		for i := len(result) - 2; i >= 0; i-- {
			result = append(result, result[i])
		}

		return strings.Join(result, "\n"), nil
	}

	return "", errors.New("wrong letter")
}

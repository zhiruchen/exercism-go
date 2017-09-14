package twelve

import (
	"fmt"
)

const testVersion = 1

var days = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var gifts = []string{
	"a Partridge in a Pear Tree.",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

func Song() string {
	s := ""
	for i := 1; i <= 12; i++ {
		s += Verse(i) + "\n"
	}
	return s
}

func Verse(day int) string {
	start := `On the %s day of Christmas my true love gave to me, `
	start = fmt.Sprintf(start, days[day-1])

	if day == 1 {
		return start + gifts[0]
	}

	for i := day; i > 1; i-- {
		start += gifts[i-1] + ", "
	}

	start += "and " + gifts[0]
	return start
}

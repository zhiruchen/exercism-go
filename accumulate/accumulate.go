package accumulate

const testVersion = 1

func Accumulate(xs []string, fn func(string) string) []string {
	newXs := []string{}

	for _, x := range xs {
		newXs = append(newXs, fn(x))
	}

	return newXs
}

package summultiples

const testVersion = 2

func SumMultiples(limit int, divisors ...int) int {
	var sum int
	var m = make(map[int]bool)
	for i := 1; i < limit; i++ {
		for _, n := range divisors {
			if i%n == 0 {
				m[i] = true
			}
		}
	}
	for k := range m {
		sum += k
	}
	return sum
}

package diffsquares

const testVersion = 1

func SquareOfSums(n int) int {
	sum := 0

	for j := 1; j <= n; j++ {
		sum += j
	}

	return sum * sum
}

func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}

	return sum
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

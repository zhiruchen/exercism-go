package pascal

func Triangle(n int) [][]int {
	start := [][]int{{1}}

	for i := 2; i <= n; i++ {
		inner := []int{1}
		prev := start[i-1]

		for j := 1; j < i-1; j++ {
			inner = append(inner, prev[j]+prev[j-1])
		}
		inner = append(inner, 1)

		start = append(start, inner)
	}
	return start
}

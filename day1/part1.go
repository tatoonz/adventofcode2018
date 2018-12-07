package day1

func part1(n []int) int {
	result := 0

	for _, nn := range n {
		result += nn
	}

	return result
}

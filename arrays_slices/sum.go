package arraysslices

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(slices ...[]int) []int {
	sums := make([]int, len(slices))

	for i, slice := range slices {
		total := 0
		for _, number := range slice {
			total += number
		}
		sums[i] = total
	}

	return sums
}

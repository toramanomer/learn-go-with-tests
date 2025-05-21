package sum

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	sums := make([]int, len(numbers))
	for i, nums := range numbers {
		sums[i] = Sum(nums)
	}
	return sums
}

func SumAllTails(numbers ...[]int) []int {
	sums := make([]int, len(numbers))

	for i, nums := range numbers {
		if len(nums) > 0 {
			sums[i] = Sum(nums[1:])
		}
	}
	return sums
}

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	result := make([]int, 2)
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			result[0] = left + 1
			result[1] = right + 1
			break
		} else if sum < target {
			left ++
		} else {
			right--
		}
	}
	return result
}

func findMin(nums []int) int {
	n := len(nums)

	// // The section might be already sorted
	// if nums[0] < nums[n-1] {
	// 	return nums[0]
	// }
	if n == 2 {
		if nums[0] < nums[1] {
			return nums[0]
		}
		return nums[1]
	} else if n == 1 {
		return nums[0]
	} else if n == 0 {
		return 0
	}
	mid := n / 2


	left := findMin(nums[:mid])
	right := findMin(nums[mid:])
	if left < right {
		return left
	}
	return right
}

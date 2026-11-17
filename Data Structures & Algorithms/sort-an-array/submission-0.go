func sortArray(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	if n == 2 {
		if nums[0] > nums[1] {
			tmp := nums[1]
			nums[1] = nums[0]
			nums[0] = tmp
		}
		return nums
	}

	midpoint := n/2
	
	
	left := sortArray(nums[:midpoint])
	right := sortArray(nums[midpoint:])
	// merge the results
	result := make([]int, n)
	l, r :=0, 0
	for l + r < n {
		if l >= len(left) {
			// take from right	
			result[l+r] = right[r]
			r++
		} else if r >= len(right) || left[l] <= right[r] {
			// take from left
			result[l + r] = left[l]
			l++
		} else {
			// take from right
			result[l+r] = right[r]
			r++
		}
	}
	return result
}

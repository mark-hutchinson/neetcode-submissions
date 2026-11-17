func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left :=0
	right := len(nums) -1
	for left <= right {
		mid := (right - left) / 2 + left
		if nums[mid] == target {
			return mid // found it
		} else if target < nums[mid]  {
			// search left of midpoint
			right = mid -1
		} else {
			// search right side
			left = mid + 1
		}
	}
	return -1
}

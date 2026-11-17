func search(nums []int, target int) int {
	n := len(nums)

	// base case: 1 element, its either the target or it isnt
	if n == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	// If the array is sorted, AND the target is not in it. we can stop searching
	if nums[0] < nums[n-1] && (target < nums[0] || target > nums[n-1]){
		return -1
	}

	// Binary search
	mid := n/2
	leftResult := search(nums[:mid], target)
	rightResult := search(nums[mid:], target)
	if leftResult > -1 {
		return leftResult
	}
	if rightResult > -1 {
		return mid + rightResult
	}
	return -1
}

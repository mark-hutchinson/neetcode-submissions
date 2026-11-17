func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}
	// note: modifies the input slice
	sort.Ints(nums)
	var result [][]int
	
	for a:=0; a < n-2; a++ {
		for a > 0 && a < n -2 && nums[a] == nums[a-1] {
			a++ // skip over previously seen a values
		}
		b := a + 1
		c := n -1
		for b < c {
			sum := nums[a] + nums[b] + nums[c]
			if sum == 0 {
				// we found a solution
				result = append(result, []int{nums[a], nums[b], nums[c]})
				// reset left pointer, making sure not to use same value again
				b++
				for nums[b] == nums[b-1] && b < c {
					b++
				}
				
				// don't need to reset right pointer. because left one is getting bigger

			} else if sum < 0 {
				// we need to increase the number. increment left pointer
				b++
			} else {
				// we need to decrease the number. decrement right pointer
				c--
			}
		}
	}
	return result
}


func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}
	sort.Ints(nums)

	var result [][]int

	for a :=0; a < n -2; a++ {
		b := a+1
		c := n-1
		if a > 0 && nums[a] == nums[a-1] {
			continue // don't consider the same A we've already looked ad
		}

		for b < c {
			sum := nums[a] + nums[b] + nums[c]
			if sum == 0 {
				result = append(result, []int{nums[a], nums[b], nums[c]})
				b++
				for nums[b] == nums[b-1] && b < c{
					b++
				}
			} else if sum > 0 {
				// too big, move the right pointer lower
				c--
			} else if sum < 0 {
				// to small, increase the number by moving left pointer right
				b++
			}
		}
	}
	return result
}


func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Printf("sorted %v\n", nums)

	if len(nums) < 3 {
		return nil
	}
	a := 0
	b := 1
	c := len(nums) - 1

	var result [][]int
	for a < len(nums) - 2 {

		for b < c {
			sum := nums[a] + nums[b] + nums[c]
			if sum == 0 {
				// found a solution
				result = append(result, []int{nums[a], nums[b], nums[c]})
				// there might be more solutions

				for b < len(nums) -1 {
					b++
					if nums[b] != nums[b-1] {
						break
					}
				}
			
				c = len(nums) - 1
			}
			if sum > 0 {
				// need to decrease the number, so decrement the c pointer
				c--
			}
			if sum < 0 {
				// need to increase the number, increment b pointer
				b++
			}
		}

		didIncrement := false
		for a < len(nums) - 2 {
			a++ // increment a pointer until it points to a different number
			if nums[a] != nums[a-1] {
				b = a + 1 // b is next element
				c = len(nums) - 1 // c is last element
				didIncrement = true
				break
			}
		}
		if !didIncrement {
			break
		}
	}

	return result
}

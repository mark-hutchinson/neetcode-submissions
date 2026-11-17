func productExceptSelf(nums []int) []int {
	totalProduct := 1
	zeroCount := 0
	for _, num := range nums {
		totalProduct = totalProduct * num
		if num == 0 {
			zeroCount++
		}
	}
	result := make([]int, len(nums))
	for i, num := range nums {
		if totalProduct != 0 {
			result[i] = totalProduct/num
		} else if totalProduct == 0 {
			if num != 0 || zeroCount > 1{
				// a zero in any of the multiples means product is 0
				result[i] = 0
			} else {
				// this is the only zero, need to exclude it
				product := 1
				for j:=0; j<len(nums); j++ {
					if i==j {
						continue
					}
					product = product * nums[j]
				}
				result[i] = product
			}
		}
	}
	return result
}

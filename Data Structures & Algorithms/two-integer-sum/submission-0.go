func twoSum(nums []int, target int) []int {
	numToIndices := make(map[int][]int)
	for i, n := range nums {
		numToIndices[n] = append(numToIndices[n], i)
	}
	
	for i, n := range nums {
		searchFor := target - n
		
		indices, ok := numToIndices[searchFor]
		if !ok {
			continue
		}
		for _, j := range indices {
			if i != j {
				// we found a solution
				return []int{i, j} // i will be lower than j
			}
		}
	}

    return nil
}

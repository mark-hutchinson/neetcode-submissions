func twoSum(nums []int, target int) []int {
	seenIndex := make(map[int]int)
	for i, n := range nums {
		needed := target - n

		// have we already seen the number we need?
		if j, ok := seenIndex[needed]; ok {
			// found the solution. 
			// j will be lower since already seen
			return []int{j, i}
		}
		seenIndex[n] = i
	}
	return nil
}

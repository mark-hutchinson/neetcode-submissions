func twoSum(nums []int, target int) []int {
	seenIndex := make(map[int]int)
	for i, n := range nums {
		needed := target - n

		// have we already seen the number we need?
		j, ok := seenIndex[needed]
		if ok {
			// found the solution. J will be lower since already seen
			return []int{j, i}
		}
		seenIndex[n] = i
	}
	return nil
}

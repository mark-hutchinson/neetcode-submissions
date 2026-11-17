func longestConsecutive(nums []int) int {

	// First, construct a map for quick lookups
	numMap := make(map[int]bool)
	for _, n := range nums {
		numMap[n] = true
	}

	longestSequence := 0
	// Iterate through the numbers. If it could be the start of the sequence
	// then count how long it could be
	for _, n := range nums {
		_, ok := numMap[n - 1]
		if ok {
			// n-1 exists. this can not be start of sequence
			continue
		}

		// This could be the start of a sequence
		sequenceLen := 1
		for i:=1;;i++ {
			nextNum := n + i
			_, ok := numMap[nextNum]
			if ok {
				// The number exists
				sequenceLen++
				continue
			}
			// the number does not exist. See if this is a new max
			if sequenceLen > longestSequence {
				longestSequence = sequenceLen
			}
			break; // end of sequence
		}
	}
	return longestSequence
}

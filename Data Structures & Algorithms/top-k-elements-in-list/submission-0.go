
func topKFrequent(nums []int, k int) []int {

	// keep track of frequency in a map num->count
	numCounts := make(map[int]int)
	for _, n := range nums {
		numCounts[n]++
	}

	// Create an array where the index is count, and each element is a slice of numbers with that count
	countToNums := make([][]int, len(nums) + 1)
	for num, count := range numCounts {
		countToNums[count] = append(countToNums[count], num)
	}

	// Iterate backwards through the array
	result := make([]int, 0, k)
	for i := len(countToNums) -1; i >0; i-- {
//		fmt.Printf("numbers with count %d: %v\n", i, countToNums[i])
		for _, num := range countToNums[i] {
			result = append(result, num)
			if len(result) >= k {
				return result
			}
		}
	}

	return result
}

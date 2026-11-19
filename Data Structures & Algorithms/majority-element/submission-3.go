func majorityElement(nums []int) int {
    counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	
	majority := (len(nums)+1)/2
	for num, count := range counts {
		if count >= majority {
			return num
		}
	}
	return 0
}

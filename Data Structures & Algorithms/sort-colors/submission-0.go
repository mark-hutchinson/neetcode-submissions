func sortColors(nums []int) {
    counts := make([]int, 3)
	for _, n := range nums {
		counts[n]++
	}
	i := 0
	for n, count := range counts {
		for j:=0; j<count;j++ {
			nums[i] = n
			i++
		}
	}
}

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := nums[0]
	for i:=1; i<len(nums); i++ {
		n = n ^ nums[i]
	}
	return n
}

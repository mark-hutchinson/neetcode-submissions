func productExceptSelf(nums []int) []int {
	count := len(nums)
	prefix := make([]int, count)
	postfix := make([]int, count)
	for i := 0; i < count; i++ {
		if i == 0 {
			prefix[i] = nums[i]
			postfix[count-1] = nums[count-1]
			continue
		}
		prefix[i] = prefix[i - 1] * nums[i]
		postfix[count - 1 -i] = postfix[count-i] * nums[count-1-i]
	}

	result := make([]int, count)
	for i:=0; i<count; i++ {
		pre :=1
		post := 1
		if i > 0 {
			pre = prefix[i - 1]
		}
		if i < count - 1 {
			post = postfix[i+1]
		}
		result[i] = pre * post
	}
	return result
}

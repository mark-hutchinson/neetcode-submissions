func climbStairs(n int) int {
    if n <= 2 {
		return n
	}
	prev := 1
	curr := 2
	for i:=2; i<n; i++ {
		tmp := curr
		curr = prev + curr
		prev = tmp
	}

	// It's fibonacci
	// 1: 1,  (1)
	// 2: 11, 2 (2)
	// 3: 111, 21, 12 (3)
	// 4: 1111, 211, 22, 121, 112 (5)
	// 5: 11111, 2111, 221, 122, 212, 1211, 1121, 1112 (8)
	return curr
}

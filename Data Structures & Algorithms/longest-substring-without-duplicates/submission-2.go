func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int)
	n := len(s)
	if n < 2 {
		return n
	}

	longest := 0
	left := 0
	right := 1
	seen[s[0]]++
	for right < n {
		seenCount := seen[s[right]]
		//fmt.Printf("Right=%d, seenCount=%d\n", right, seenCount)
		if seenCount > 0 {
			// shrink the window on the left
			seen[s[left]]--
			left++
		} else {
			seen[s[right]]++
			l := right - left + 1
			if l > longest {
				longest = l
			}
			// increase the window on the right
			right++
		}
	}
	return longest

}

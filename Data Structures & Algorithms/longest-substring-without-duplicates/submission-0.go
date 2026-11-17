func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]struct{})
	n := len(s)
	if n < 2 {
		return n
	}

	longest := 0
	left := 0
	right := 1
	seen[s[0]] = struct{}{}
	for right < n {
		_, duplicate := seen[s[right]]
		seen[s[right]] = struct{}{} 
		if duplicate {
			// shrink the window on the left
			delete(seen, s[left])
			left++
		} else {
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


func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	
	targetChars := make(map[byte]int)
	for _, r := range s1 {
		targetChars[byte(r)]++
	}

	windowChars := make(map[byte]int)
	left := 0
	right := len(s1) - 1
	for i :=0; i<=right - 1; i++ {
		windowChars[s2[i]]++
	}
	for right < len(s2) {
		// add right to window
		windowChars[s2[right]]++

		if isAnagram(targetChars, windowChars) {
			return true
		}
		// remove left from the window
		windowChars[s2[left]]--
		left++
		right++

	}

	return false
}

func isAnagram(m1 map[byte]int, m2 map[byte]int) bool{
	// log.Printf("comparing %v, %v\n", m1, m2)
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}

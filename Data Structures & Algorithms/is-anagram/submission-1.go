func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// keep track of the count of each char in the string s
	charCount := make(map[rune]int)
	for _, c := range s {
		charCount[c]++
	}

	// for string t, see if there is a char count in the map for each char encountered
	for _, c := range t {
		if charCount[c] == 0 {
			return false 
		}
		charCount[c]--
	}
	return true
}

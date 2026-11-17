func isAnagram(s string, t string) bool {

	charCount := make(map[rune]int)
	// keep track of the count of each char in the string s
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

	for _, count := range charCount {
		if count > 0 {
			return false // there were chars in S that were not used in T
		}
	}
	return true

}

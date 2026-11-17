func canConstruct(ransomNote string, magazine string) bool {

	availableChars := make(map[rune]int)
	for _, r := range magazine {
		availableChars[r]++
	}

	for _, r := range ransomNote {
		available := availableChars[r]
		if available == 0 {
			return false
		}
		availableChars[r]--
	}
	return true
	
}

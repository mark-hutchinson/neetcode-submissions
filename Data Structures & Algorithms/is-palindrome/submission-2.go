func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	
	left := 0
	right := len(s) - 1
	for left < right {
		
		// skip non alphanumeric chars
		for left < right && !isAlphaNumeric(s[left]) {
			left++
		}
		for left < right && !isAlphaNumeric(s[right]) {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphaNumeric(char byte) bool {
	if char >= 'a' && char <= 'z'{
		return true
	}
	if char >= '0' && char <= '9'{
		return true
	}
	return false
}

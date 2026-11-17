func removeVowels(s string) string {
    var sb strings.Builder

	for _, b := range s {
		if !isVowel(b) {
			sb.WriteRune(b)
		}
	}
	return sb.String()
}

func isVowel(b rune) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}

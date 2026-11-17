
func groupAnagrams(strs []string) [][]string {
	// Do one pass through the strings, computing the counts
	// e.g. aaab - computes to 3,1,0,0,0..
	letterCounts := make([][]int, len(strs))
	for i, str := range strs {
		letterCounts[i] = countLetters(str)
	}

	// Then do a grouping matching counts in a map
	// key: string value of letter frequency. value: matching words
	resultMap := make(map[string][]string)
	for i, counts := range letterCounts {
		strCounts := fmt.Sprintf("%v", counts)
		resultMap[strCounts] = append(resultMap[strCounts], strs[i])
	}

	// Get all the values from the map. this is the result
	result := make([][]string, 0)
	for _, val := range resultMap {
		result = append(result, val)
	}
	return result
}

func countLetters(str string)[]int {
	counts := make([]int, 26)
	for _, c := range str {
		index := c - 'a'
		counts[index]++
	}	
	return counts
}
func characterReplacement(s string, k int) int {


    // Idea: sliding window. left pointer at start of window, right pointer at right
    // keep track of char counts in that window
    n := len(s)

    if k >= n - 1 {
        // we can replace all but one elements. the whole input can be made consecutive
        return n
    }

    charCount := make(map[byte]int)
    left := 0
    right := 1
    longest :=1
    charCount[s[left]]++

    // also, we will keep track of the most frequent letter observed in the window. This is the one we will not need to replace
    mostFrequentChar := s[left]
    mostFrequentCount := 1

    for right < n {

        charCount[s[right]]++
        if charCount[s[right]] > mostFrequentCount {
            // we have a new most frequent char in the window
            mostFrequentChar = s[right]
            mostFrequentCount = charCount[s[right]]
        }

        length := right - left + 1
        numReplacementsNeeded := length - mostFrequentCount
        if numReplacementsNeeded > k {
            // slide the window by moving left pointer
            charCount[s[left]]--
            if mostFrequentChar == s[left] {
                // potentially a new most frequent char. find it
                mostFrequentCount = 0
                for c, count := range charCount {
                    if count > mostFrequentCount {
                        mostFrequentCount = count
                        mostFrequentChar = c
                    }
                }
            }
            left++
            right++ // we never need to shrink the window
            //fmt.Printf("sliding window left=%d, right=%d\n", left, right)
        } else {
            // this is a valid candidate
            if length > longest {
                longest = length
            }
            // Increase the window size
            right ++
            //fmt.Printf("increasing window. left=%d, right=%d\n", left, right)
        }
    }
    return longest
}

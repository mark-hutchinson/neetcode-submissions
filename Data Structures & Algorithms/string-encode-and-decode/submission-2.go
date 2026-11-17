
type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    // Protocol:
    // {numStrings},{len0},{len1},strs[0]strs[1]
    header := make([]string, len(strs) +1)
    header[0] = fmt.Sprintf("%d", len(strs))
    for i, str := range strs{
        header[i+1] = fmt.Sprintf("%d", len(str))
    }
    
    var builder strings.Builder
    for _, headerVal := range header {
        builder.WriteString(headerVal)
        builder.WriteString(",")
    }
    for _, str := range strs {
        builder.WriteString(str)
    }
    encoded := builder.String()
    return encoded
}

func (s *Solution) Decode(encoded string) []string {
    numWords, remaining := popNumber(encoded)

    wordLens := make([]int, numWords)
    for i:=0; i<numWords; i++ {
        var wordLen int
        wordLen, remaining = popNumber(remaining)
        wordLens[i] = wordLen
    }

    result := make([]string, numWords)
    for i:=0; i<numWords; i++ {
        wordLen := wordLens[i]
        var word string
        word, remaining = popLen(remaining, wordLen)
        result[i] = word
    }
    return result
}


func popNumber(encoded string) (int, string) {
    commaPos := strings.Index(encoded,",")
    numberStr := encoded[0:commaPos]
    remaining := encoded[commaPos+1:]
    numberVal, err := strconv.Atoi(numberStr)
    if err != nil {
        panic("could not parse num")
    }
    return numberVal, remaining
}

func popLen(encoded string, length int) (string, string) {
    popped := encoded[:length]
    remaining := encoded[length:]
    return popped, remaining
}


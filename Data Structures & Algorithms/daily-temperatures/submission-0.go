
type Tuple struct {
	Index int
	Value int
}


func dailyTemperatures(temperatures []int) []int {
	var stack []Tuple
	var result []int = make([]int, len(temperatures))
	for i, t := range temperatures {
		for len(stack) > 0 && stack[len(stack) -1].Value < t {
			// this temperature is greater than a previous day
			popped := stack[len(stack)-1]
			stack = stack [:len(stack)-1]
			result[popped.Index] = i - popped.Index
		}

		// Push onto the stack. It will be popped if we find a future day
		stack = append(stack, Tuple{Index:i, Value:t})
	}

	return result
}

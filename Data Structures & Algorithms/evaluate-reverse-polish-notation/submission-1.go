func evalRPN(tokens []string) int {
	var stack []int

	for _, t := range tokens {
		numVal, err := strconv.Atoi(t)
		if err != nil{
			// We expect an operation
			n := len(stack)
			if n < 2 {
				panic("got an operation but not enough numbers on stack")
			}
			// pop a
			a := stack[n-1]
			// pop b
			b:= stack[n-2]
			stack = stack[:n-2]
			// evaluate
			result := 0
			switch t {
				case "+":
					result = a + b
				case "-":
					result = b - a
				case "*":
					result = a * b
				case "/":
					result = b / a
				default:
					panic(fmt.Sprintf("invalid operation %s", t))
			}
			// push result to stack
			stack = append(stack, result)
		} else {
			// we have a number, push to the stack
			stack = append(stack, numVal)
		}
	}

	return stack[0]
}




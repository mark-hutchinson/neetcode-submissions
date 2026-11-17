type Stack struct {
	bytes []byte
}


func isValid(s string) bool {
	stack := &Stack{}
	for _, b := range []byte(s) {
		if isOpen(b) {
			stack.push(b)
		} else if isClose(b) {
			popped, err := stack.pop()
			if err != nil {
				return false
			}
			// check if the popped matches
			if !matchesClosing(popped, b) {
				return false
			}
		} else {
			fmt.Printf("Invalid char encountered %c\n", b)
			return false
		}
	}
	return stack.isEmpty()
}

func isOpen(b byte) bool {
	return b == '(' || b == '{' || b == '['
}

func isClose(b byte) bool {
	return b == ')' || b == '}' || b == ']'
}

// matchesClosing returns true if the open bracket matches the closing bracket
func matchesClosing(open, closing byte) bool {
	switch open {
		case '(':
			return closing == ')'
		case '{':
			return closing == '}'
		case '[':
			return closing == ']'
		default:
			fmt.Printf("invalid open brace %c\n", open)
			return false
	}
}


func (s *Stack) push(b byte) {
	s.bytes = append(s.bytes, b)
}

func (s *Stack) pop() (byte, error) {
	n := len(s.bytes)
	if n == 0 {
		return 0, fmt.Errorf("can't pop empty stack")
	}
	val := s.bytes[n-1]
	s.bytes = s.bytes[:n-1]
	return val, nil
}

func (s *Stack) isEmpty() bool {
	return len(s.bytes) == 0
}
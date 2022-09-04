package stack

// To evaluate if an expression is balanced
// expr? {[()]}

func IsBalancedExpr(expr string) bool {
	stack := NewStack[int32]()

	for _, char := range expr {
		switch char {
		case '{':
			stack.Push('}')
			break
		case '[':
			stack.Push(']')
			break
		case '(':
			stack.Push(')')
			break
		default:
			if item, err := stack.Pop(); err != nil || item != char {
				return false
			}

		}
	}

	if !stack.Empty() {
		return false
	}

	return true
}

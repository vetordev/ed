package stack

// To evaluate if an expression is balanced
// expr? {[()]}

func IsBalancedExpr(expr string) bool {
	stack := NewStack[int32]()

	for _, char := range expr {
		switch char {
		case '{':
			stack.push('}')
			break
		case '[':
			stack.push(']')
			break
		case '(':
			stack.push(')')
			break
		default:
			if item, err := stack.pop(); err != nil || item != char {
				return false
			}

		}
	}

	if !stack.empty() {
		return false
	}

	return true
}

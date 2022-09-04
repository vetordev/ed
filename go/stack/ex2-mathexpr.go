package stack

import (
	"strconv"
	"strings"
	"unicode"
)

// expr: ((2 * 3) + (3 / 2)) -> 23*32/+
func PostfixMathExpr(expr string) string {
	stack := NewStack[int32]()
	postfixExpr := ""

	for _, char := range expr {
		if char == ')' {
			operator, _ := stack.Pop()

			postfixExpr += string(operator)
		}

		if unicode.IsNumber(char) {
			postfixExpr += string(char)
		}

		if strings.Contains("+-*/", string(char)) {
			stack.Push(char)
		}

	}

	return postfixExpr
}

func SumPostFixedExpr(expr string) float64 {

	stack := NewStack[float64]()

	for _, char := range expr {
		if conv, err := strconv.ParseFloat(string(char), 64); err == nil {
			stack.Push(conv)
			continue
		}

		value2, _ := stack.Pop()
		value1, _ := stack.Pop()

		switch char {
		case '+':
			stack.Push(value1 + value2)
			break
		case '-':
			stack.Push(value1 - value2)
			break
		case '*':
			stack.Push(value1 * value2)
			break
		case '/':
			stack.Push(value1 / value2)
			break
		}
	}

	r, _ := stack.Pop()

	return r
}

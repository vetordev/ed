package stack

import (
	"strconv"
	"strings"
	"unicode"
)

// expr: ((2 * 3) + (3 / 2)) -> 23*32/+
func PostfixMathExpr(expr string) (string, error) {
	stack := NewStack[int32]()
	postfixExpr := ""

	for _, char := range expr {
		if char == ')' {
			operator, _ := stack.pop()

			postfixExpr += string(operator)
		}

		if unicode.IsNumber(char) {
			postfixExpr += string(char)
		}

		if strings.Contains("+-*/", string(char)) {
			stack.push(char)
		}

	}

	return postfixExpr, nil
}

func SumPostFixedExpr(expr string) float64 {

	stack := NewStack[float64]()

	for _, char := range expr {
		if conv, err := strconv.ParseFloat(string(char), 64); err == nil {
			stack.push(conv)
			continue
		}

		value2, _ := stack.pop()
		value1, _ := stack.pop()

		switch char {
		case '+':
			stack.push(value1 + value2)
			break
		case '-':
			stack.push(value1 - value2)
			break
		case '*':
			stack.push(value1 * value2)
			break
		case '/':
			stack.push(value1 / value2)
			break
		}
	}

	r, _ := stack.pop()

	return r
}

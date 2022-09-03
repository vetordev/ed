package stack

import (
	"testing"
)

func TestBalanced(t *testing.T) {
	expr := "{([])}"
	balanced := IsBalancedExpr(expr)

	want := true

	if balanced != want {
		t.Fatalf(`Balanced(%q) = %v, want match for %v`, expr, balanced, want)
	}
}

func TestNotBalanced(t *testing.T) {
	expr := "{([})"
	balanced := IsBalancedExpr(expr)

	want := false

	if balanced != want {
		t.Fatalf(`Balanced(%q) = %v, want match for %v`, expr, balanced, want)
	}
}

func TestPostFix(t *testing.T) {
	mathExpr := "((2*3)+(8/4))"
	postfix := PostfixMathExpr(mathExpr)

	want := "23*84/+"

	if postfix != want {
		t.Fatalf(`PostfixMathExpr(%q) = %q, want match for %q`, mathExpr, postfix, want)
	}
}

func TestSumPostFixed(t *testing.T) {
	postFixed := "23*84/+"
	sum := SumPostFixedExpr(postFixed)

	want := 8.0

	if sum != want {
		t.Fatalf(`SumPostFixedExpr(%q) = %g, want match for %g`, postFixed, sum, want)
	}
}

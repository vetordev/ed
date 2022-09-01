package stack

import (
	"testing"
)

func TestBalanced(t *testing.T) {
	balanced := IsBalancedExpr("{([])}")

	if balanced != true {
		t.Fatalf(`Balanced("{([])}") = %q`, "a")
	}
}

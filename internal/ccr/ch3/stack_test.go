package stack

import (
	"testing"
)

func Test(t *testing.T) {
	s1 := New()
	s1.Push(1)
	s1.Push(2)
	if s1.Pop() != 2 {
		t.Errorf("Top item on the stack should be 2")

	}
}

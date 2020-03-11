package chapter1

import "testing"

func TestIsUnique(t *testing.T) {
	tests := []struct {
		str string

		ans bool
	}{
		{"abcd", true},
		{"aabc", false},
		{"", true}}

	for _, test := range tests {
		if isUnique(test.str) != test.ans {
			t.Errorf("IsUnique Test failed for %s", test.str)
		}
	}

}

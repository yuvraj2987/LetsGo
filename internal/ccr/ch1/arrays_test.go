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

func TestCheckPermutation(t *testing.T) {
	tests := []struct {
		s1  string
		s2  string
		ans bool
	}{
		{"a", "a", true},
		{"ab", "ba", true},
		{"abcd", "efgh", false},
		{"abba", "bbaa", true},
		{"abc", "abcd", false}}

	for _, test := range tests {
		if checkPermutation(test.s1, test.s2) != test.ans {
			t.Errorf("IsUnique Test failed for s1=%s and s2=%s", test.s1, test.s2)
		}
	}

}

func TestIsPalinPermutation(t *testing.T) {
	tests := []struct {
		s   string
		ans bool
	}{
		{"", false},
		{"abc", false},
		{"Tact Coa", true},
		{"aba", true}}

	for _, test := range tests {
		if IsPalinPermutation(test.s) != test.ans {
			t.Errorf("Test failed for string %s", test.s)
		}
	}

}

/**

func TestUrlify(t *testing.T) {
	r1 := []rune("Hi John")

	tests := []struct {
		s1 []rune
		s2 []rune
	}{
		{},
		{"abc", "abcd", false}}

	for _, test := range tests {
		if checkPermutation(test.s1, test.s2) != test.ans {
			t.Errorf("IsUnique Test failed for %s", test.str)
		}
	}

}
**/

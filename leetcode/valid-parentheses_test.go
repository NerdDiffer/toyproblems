package leetcode

import "testing"

// The brackets must close in the correct order, "()" and "()[]{}" are all valid
// but "(]" and "([)]" are not.
func TestIsValid(t *testing.T) {
	tables := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},

		{"[", false},
		{"]", false},
	}

	for _, test := range tables {
		actual := IsValid(test.input)

		if actual != test.expected {
			t.Errorf("\ninput: %s\nactual: %t\nexpected: %t", test.input, actual, test.expected)
		}
	}
}

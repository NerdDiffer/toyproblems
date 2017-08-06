package leetcode

import "testing"

func TestReverseStr(t *testing.T) {
	tables := []struct {
		s        string
		k        int
		expected string
	}{
		{"racecar", 7, "racecar"},
		{"foobar", 6, "raboof"},
		{"abcdefg", 2, "bacdfeg"},
		{"abcdefg", 1, "abcdefg"},
		{"abcd", 2, "bacd"},
		{"abcdefg", 4, "dcbaefg"},
	}

	for _, test := range tables {
		actual := ReverseStr(test.s, test.k)

		if actual != test.expected {
			t.Errorf("actual: %s, expected: %s", actual, test.expected)
		}
	}
}

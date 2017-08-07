package leetcode

import "testing"

func TestSingleNumber(t *testing.T) {
	tables := []struct {
		input    []int
		expected int
	}{
		{input: []int{1, 2, 3, 2, 1}, expected: 3},
	}

	for _, test := range tables {
		actual := SingleNumber(test.input)
		if actual != test.expected {
			t.Errorf("\ninput: %v\nactual: %v\nexpected: %v", test.input, actual, test.expected)
		}
	}
}

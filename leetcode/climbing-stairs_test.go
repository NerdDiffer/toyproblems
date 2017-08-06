package leetcode

import "testing"

func TestClimbStairs(t *testing.T) {
	tables := []struct {
		input    int
		expected int
	}{
		{3, 3},
		{4, 5},
		{5, 8},
	}

	for _, test := range tables {
		actual := ClimbStairs(test.input)

		if actual != test.expected {
			t.Errorf("actual: %d, expected: %d", actual, test.expected)
		}
	}
}

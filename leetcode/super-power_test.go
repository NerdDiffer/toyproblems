package leetcode

import "testing"

func TestSuperPow(t *testing.T) {
	tables := []struct {
		a        int
		b        []int
		expected int
	}{
		{2, []int{3}, 8},
		{2, []int{1, 0}, 1024},
		{2, []int{0}, 1},
		{2147483647, []int{2, 0, 0}, 1198},
	}

	for _, test := range tables {
		actual := SuperPow(test.a, test.b)

		if actual != test.expected {
			t.Errorf("\ninputs:   (%v, %v)\nactual:   %d \nexpected: %d", test.a, test.b, actual, test.expected)
		}
	}
}

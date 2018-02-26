package leetcode

import "testing"

func TestArrayPartitionI(t *testing.T) {
	tables := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 4, 3, 2}, 4},
		{[]int{-5, 4, -1, 2, -3, 6}, -2},
	}
	// (-5,-3) (-1, 2) (4,6) // -5 - 1 + 4
	for _, test := range tables {
		actual := arrayPairSum(test.input)

		if actual != test.expected {
			t.Errorf("\ninput: %v\nactual: %d\nexpected: %d", test.input, actual, test.expected)
		}
	}
}

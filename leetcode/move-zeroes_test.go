package leetcode

import (
	tph "github.com/NerdDiffer/toyproblems/helpers"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tables := []struct {
		input    []int
		expected []int
	}{
		{input: []int{0, 1, 0, 3, 12}, expected: []int{1, 3, 12, 0, 0}},
		{input: []int{0, 0, 1}, expected: []int{1, 0, 0}},
		{input: []int{0, 2, 1, 0}, expected: []int{2, 1, 0, 0}},
	}

	for _, test := range tables {
		actual := make([]int, len(test.input))
		copy(actual, test.input)

		MoveZeroes(actual)

		if !tph.CmpIntSlices(actual, test.expected) {
			t.Errorf("\ninput: %v\nactual: %v\nexpected: %v", test.input, actual, test.expected)
		}
	}
}

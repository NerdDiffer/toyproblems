package leetcode

import (
	tph "github.com/NerdDiffer/toyproblems/helpers"
	"testing"
)

func TestFindErrorNums(t *testing.T) {
	tables := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 2, 4}, expected: []int{2, 3}},
		{input: []int{2, 2}, expected: []int{2, 1}},
		{input: []int{3, 2, 3, 4, 6, 5}, expected: []int{3, 1}},
	}

	for _, test := range tables {
		actual := FindErrorNums(test.input)
		if !tph.CmpIntSlices(actual, test.expected) {
			t.Errorf("\ninput: %v\nactual: %v\nexpected: %v", test.input, actual, test.expected)
		}
	}
}

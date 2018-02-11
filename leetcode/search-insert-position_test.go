package leetcode

import (
	"fmt"
	"testing"
)

type Input struct {
	list []int
	val  int
}

func (i Input) show() string {
	return fmt.Sprintf("(%v, %d)", i.list, i.val)
}

func TestSearchInsertPosition(t *testing.T) {
	tables := []struct {
		input    Input
		expected int
	}{
		{input: Input{list: []int{1, 3, 5, 6}, val: 5}, expected: 2},
		{input: Input{list: []int{1, 3, 5, 6}, val: 2}, expected: 1},
		{input: Input{list: []int{1, 3, 5, 6}, val: 7}, expected: 4},
		{input: Input{list: []int{1, 3, 5, 6}, val: 0}, expected: 0},
		{input: Input{list: []int{1}, val: 1}, expected: 0},
	}

	for _, test := range tables {
		list := test.input.list
		val := test.input.val

		actual := SearchInsertPosition(list, val)
		expected := test.expected

		if actual != expected {
			ppInput := test.input.show()
			t.Errorf("\ninput: %v\nactual: %v\nexpected: %v", ppInput, actual, test.expected)
		}
	}
}

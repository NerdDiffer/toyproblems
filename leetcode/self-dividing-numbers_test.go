package leetcode

import (
	"testing"
)

func TestIsSelfDivisible(t *testing.T) {
	tables := []struct {
		input    int
		expected bool
	}{
		{128, true},
		{101, false},
	}

	for _, test := range tables {
		actual := isSelfDivisible(test.input)

		if actual != test.expected {
			t.Errorf("\ninput: %d\nactual: %t\nexpected: %t\n", test.input, actual, test.expected)
		}
	}
}

func TestGetDigit(t *testing.T) {
	tables := []struct {
		inputNum   int
		inputPlace int
		expected   int
	}{
		{12345, 0, 5},
		{12345, 1, 4},
		{12345, 2, 3},
		{12345, 3, 2},
		{12345, 4, 1},
		{12345, 5, 0},
	}

	for _, test := range tables {
		actual := digitAtPlace(test.inputNum, test.inputPlace)

		if actual != test.expected {
			t.Errorf("\ninputs: (%d, %d)\nactual: %d\nexpected: %d\n", test.inputNum, test.inputPlace, actual, test.expected)
		}
	}
}

func TestSelfDividingNumbers(t *testing.T) {
	tables := []struct {
		left     int
		right    int
		expected []int
	}{
		{1, 22, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}},
		{66, 708, []int{66, 77, 88, 99, 111, 112, 115, 122, 124, 126, 128, 132, 135, 144, 155, 162, 168, 175, 184, 212, 216, 222, 224, 244, 248, 264, 288, 312, 315, 324, 333, 336, 366, 384, 396, 412, 424, 432, 444, 448, 488, 515, 555, 612, 624, 636, 648, 666, 672}},
	}

	for _, test := range tables {
		actual := selfDividingNumbers(test.left, test.right)

		if len(actual) != len(test.expected) {
			t.Errorf("length of result must be the same length as expected")
			break
		}

		for i, v := range test.expected {
			if actual[i] != v {
				t.Errorf("\ninputs: (%d, %d)\nindex: %d\nactual: %d\nexpected: %d", test.left, test.right, i, actual[i], test.expected[i])
				break
			}
		}
	}
}
